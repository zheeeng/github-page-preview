package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	matchURLPattern = iota
	matchURLWithoutHostPattern
	matchRootPattern
	matchFilePattern
)

var (
	baseExp           = regexp.MustCompile(`^/[\w-~]+/[\w-~]+(/blob|/tree)?/[\w-~]+`)
	urlExp            = regexp.MustCompile(`^/(?P<user>[\w-~]+)/(?P<repo>[\w-~]+)(/blob|/tree)?/(?P<branch>[\w-~]+)(?P<path>[\w-~/]*)//(?P<folder>[\w-~/]*?)(?P<file>(/?[^/\s]+\.[^/\s]+)?$)`)
	urlWithoutHostExp = regexp.MustCompile(`^/(?P<user>[\w-~]+)/(?P<repo>[\w-~]+)(/blob|/tree)?/(?P<branch>[\w-~]+)(?P<path>[\w-~/]*)(?P<file>(/[^/\s]+\.[^/\s]+)?$)`)
	rootExp           = regexp.MustCompile(`^/$`)
	fileExp           = regexp.MustCompile(`^/(?P<file>.*)`)
)

var (
	// ErrNotRecognizeURL presents url path no patterns matching
	ErrNotRecognizeURL = errors.New("Can't recognize the path format")
)

// PathComponents interface
type PathComponents interface {
	Endpoint() string
	StaticHost() string
	GetFile() string
}

type endpointComponents struct {
	matchType int
	user      string
	repo      string
	branch    string
	path      string
	folder    string
	file      string
}

func patternMatch(endpoint string) *endpointComponents {
	endpointBytes := []byte(endpoint)

	switch true {
	case urlExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchURLPattern}).parseFrom(endpoint, urlExp)
	case urlWithoutHostExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchURLWithoutHostPattern}).parseFrom(endpoint, urlWithoutHostExp)
	case rootExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchRootPattern}).parseFrom(endpoint, rootExp)
	case fileExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchFilePattern}).parseFrom(endpoint, fileExp)
	default:
		// endpoint is empty string
		return nil
	}
}

// NewEndpointComponents returns endpointComponents instance
func NewEndpointComponents(path string, referer string) (PathComponents, error) {
	refEC := patternMatch(referer)
	pathEC := patternMatch(path)

	if pathEC == nil {
		panic("impossible situation")
	}

	if refEC == nil {
		if pathEC.matchType == matchRootPattern || pathEC.matchType == matchFilePattern {
			return nil, ErrNotRecognizeURL
		}

		return pathEC, nil
	}

	if refEC.matchType == matchRootPattern || refEC.matchType == matchFilePattern {
		return NewEndpointComponents(path, "")
	}

	return refEC.setFile(pathEC.GetFile()), nil
}

func (uc *endpointComponents) setFile(file string) *endpointComponents {
	uc.file = file
	return uc
}

func (uc *endpointComponents) GetFile() string {
	return uc.file
}

func (uc *endpointComponents) parseFrom(path string, reg *regexp.Regexp) *endpointComponents {
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

func (uc *endpointComponents) Endpoint() string {
	return fmt.Sprintf("/%s/%s/%s%s%s%s", uc.user, uc.repo, uc.branch, uc.path, uc.folder, uc.file)
}

func (uc *endpointComponents) StaticHost() string {
	return fmt.Sprintf("/%s/%s/%s%s", uc.user, uc.repo, uc.branch, uc.path)
}
