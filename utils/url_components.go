package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	baseExp           = regexp.MustCompile(`^/[\w-~]+/[\w-~]+(/blob|/tree)?/[\w-~]+`)
	urlExp            = regexp.MustCompile(`^/(?P<user>[\w-~]+)/(?P<repo>[\w-~]+)(/blob|/tree)?/(?P<branch>[\w-~]+)(?P<path>[\w-~/]*)//(?P<folder>[\w-~/]*?)(?P<file>(/?[^/\s]+\.[^/\s]+)?$)`)
	urlWithoutHostExp = regexp.MustCompile(`^/(?P<user>[\w-~]+)/(?P<repo>[\w-~]+)(/blob|/tree)?/(?P<branch>[\w-~]+)(?P<path>[\w-~/]*)(?P<file>(/[^/\s]+\.[^/\s]+)?$)`)
	fileExp           = regexp.MustCompile(`^/(?P<file>.*)`)
	// ErrNotMatchURLPattern presents url path no patterns matching
	ErrNotMatchURLPattern = errors.New("Can't recognize the path format")
)

// PathComponents interface
type PathComponents interface {
	Endpoint() string
	StaticHost() string
	setFile(string)
	GetFile() string
}

type pathComponents struct {
	user   string
	repo   string
	branch string
	path   string
	folder string
	file   string
}

// NewPathComponents returns pathComponents instance
func NewPathComponents(path string, referer string) (PathComponents, error) {
	pathBytes := []byte(path)

	switch true {
	// Situation 1: host is specified by delimiter '//'
	case urlExp.Match(pathBytes):
		return (&pathComponents{}).parseFrom(path, urlExp), nil
	// Situation 2: host are detected by default rule
	case urlWithoutHostExp.Match(pathBytes):
		return (&pathComponents{}).parseFrom(path, urlWithoutHostExp), nil
		// Situation 3: path is relative path to root, we get host by referer
	case fileExp.Match(pathBytes) && referer != "":
		pc := (&pathComponents{}).parseFrom(path, fileExp)
		refPc, err := NewPathComponents(referer, "")
		refPc.setFile(pc.GetFile())
		return refPc, err
	}

	return nil, ErrNotMatchURLPattern
}

func (uc *pathComponents) setFile(file string) {
	uc.file = file
}

func (uc *pathComponents) GetFile() string {
	return uc.file
}

func (uc *pathComponents) parseFrom(path string, reg *regexp.Regexp) *pathComponents {
	match := reg.FindStringSubmatch(path)

	for i, name := range reg.SubexpNames() {
		switch name {
		case "user":
			uc.user = match[i]
		case "repo":
			uc.repo = match[i]
		case "branch":
			uc.branch = match[i]
		case "path":
			uc.path = match[i]
			if strings.HasSuffix(uc.path, "/") {
				uc.path = uc.path[0 : len(uc.path)-1]
			}
		case "folder":
			uc.folder = match[i]
			if uc.folder != "" {
				uc.folder = "/" + uc.folder
			}
			if strings.HasSuffix(uc.folder, "/") {
				uc.folder = uc.folder[0 : len(uc.folder)-1]
			}
		case "file":
			uc.file = match[i]
			if !strings.HasPrefix(uc.file, "/") {
				uc.file = "/" + uc.file
			}
			if strings.HasSuffix(uc.file, "/") {
				uc.file += "index.html"
			} else if filepath.Ext(uc.file) == "" {
				uc.file += "/index.html"
			}
		}
	}

	return uc
}

func (uc *pathComponents) Endpoint() string {
	return fmt.Sprintf("/%s/%s/%s%s%s%s", uc.user, uc.repo, uc.branch, uc.path, uc.folder, uc.file)
}

func (uc *pathComponents) StaticHost() string {
	return fmt.Sprintf("/%s/%s/%s%s", uc.user, uc.repo, uc.branch, uc.path)
}
