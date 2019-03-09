package main

import (
	"fmt"
	"regexp"
)

var urlExp = regexp.MustCompile(`/(?P<user>[\w-~@]+)/(?P<repo>[\w-~@]+)/blob/(?P<branch>[\w-~@]+)/(?P<path>[\w-~@/]+)//(?P<asset>.+)`)

type pathComponents struct {
	user   string
	repo   string
	branch string
	path   string
	asset  string
}

func (uc *pathComponents) parseFrom(path string) *pathComponents {
	match := urlExp.FindStringSubmatch(path)

	for i, name := range urlExp.SubexpNames() {
		switch name {
		case "user":
			uc.user = match[i]
		case "repo":
			uc.repo = match[i]
		case "branch":
			uc.branch = match[i]
		case "path":
			uc.path = match[i]
		case "asset":
			uc.asset = match[i]
		}
	}

	return uc
}

func (uc *pathComponents) compileToRaw() string {
	return fmt.Sprintf("/%s/%s/blob/%s/%s//%s", uc.user, uc.repo, uc.branch, uc.path, uc.asset)
}

func (uc *pathComponents) compileToRequestPath() string {
	return fmt.Sprintf("/%s/%s/%s/%s/%s", uc.user, uc.repo, uc.branch, uc.path, uc.asset)
}

func (uc *pathComponents) compileToStaticHost() string {
	return fmt.Sprintf("/%s/%s/%s/%s", uc.user, uc.repo, uc.branch, uc.path)
}
