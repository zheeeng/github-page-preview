package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	urlExp            = regexp.MustCompile(`/(?P<user>[\w-~]+)/(?P<repo>[\w-~]+)/blob/(?P<branch>[\w-~]+)(?P<path>[\w-~/]*)//(?P<asset>.*)`)
	urlWithoutHostExp = regexp.MustCompile(`/(?P<user>[\w-~]+)/(?P<repo>[\w-~]+)/blob/(?P<branch>[\w-~]+)(?P<path>[\w-~/]*)(?P<asset>(/[^/\s]+\.[^/\s]+)?$)`)
	assetExp          = regexp.MustCompile(`/(?P<asset>.*)`)
)

// PathComponents interface
type PathComponents interface {
	RequestPath() string
	StaticHost() string
	setAsset(string)
	getAsset() string
}

type pathComponents struct {
	user   string
	repo   string
	branch string
	path   string
	asset  string
}

// NewPathComponents returns pathComponents instance
func NewPathComponents(path string, referer string) PathComponents {
	pathBytes := []byte(path)

	switch true {
	case urlExp.Match(pathBytes):
		return (&pathComponents{}).parseFrom(path, urlExp)
	case urlWithoutHostExp.Match(pathBytes):
		return (&pathComponents{}).parseFrom(path, urlWithoutHostExp)
	case assetExp.Match(pathBytes) && referer != "":
		pc := (&pathComponents{}).parseFrom(path, assetExp)
		refPc := NewPathComponents(referer, "")
		refPc.setAsset(pc.getAsset())
		return refPc
	}

	panic("Can't recognize the path format")
}

func (uc *pathComponents) setAsset(asset string) {
	uc.asset = asset
}

func (uc *pathComponents) getAsset() string {
	return uc.asset
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
		case "asset":
			uc.asset = match[i]
			if !strings.HasPrefix(uc.asset, "/") {
				uc.asset = "/" + uc.asset
			}
			if strings.HasSuffix(uc.asset, "/") {
				uc.asset += "index.html"
			} else if filepath.Ext(uc.asset) == "" {
				uc.asset += "/index.html"
			}
		}
	}

	return uc
}

func (uc *pathComponents) RequestPath() string {
	return fmt.Sprintf("/%s/%s/%s%s%s", uc.user, uc.repo, uc.branch, uc.path, uc.asset)
}

func (uc *pathComponents) StaticHost() string {
	return fmt.Sprintf("/%s/%s/%s%s", uc.user, uc.repo, uc.branch, uc.path)
}
