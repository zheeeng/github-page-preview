package utils

import (
	"fmt"
	"math"
	"path"
	"regexp"
	"strings"
)

// EndpointType indicates the type of  EndpointType
type EndpointType int

const (
	// EndpointLocalAsset presents reading static asset
	EndpointLocalAsset EndpointType = iota
	// EndpointRemoteAsset presents reading remote asset
	EndpointRemoteAsset
	// EndpointRedirect presents redirect endpoint to process it again
	EndpointRedirect
	// EndpointNotFound presents 404
	EndpointNotFound
)

type matchPattern int

const (
	matchBasePattern matchPattern = 1
	matchURLPattern               = 1 << iota
	matchURLWithoutHostPattern
	matchFilePattern
)

// count for endpoint types
var mpc = uint(math.Log2(float64(matchFilePattern)))

var (
	baseExp           = regexp.MustCompile(`^/(?P<user>[\w-~]+/)(?P<repo>[\w-~]+/)(?P<blob>blob/|tree/)(?P<branch>[\w-~:]+/)`)
	urlExp            = regexp.MustCompile(`^/(?P<user>[\w-~]+/)(?P<repo>[\w-~]+/)(?P<blob>blob/|tree/)(?P<branch>[\w-~]+/?)(?P<path>[\w-~/]*)::/(?P<folder>[\w-~/]*?)(?P<file>(/?[^/\s]+\.[^/\s]+)?$)`)
	urlWithoutHostExp = regexp.MustCompile(`^/(?P<user>[\w-~]+/)(?P<repo>[\w-~]+/)(?P<blob>blob/|tree/)(?P<branch>[\w-~]+/?)(?P<path>[\w-~/]*)(?P<file>(/[^/\s]+\.[^/\s]+)?$)`)
	fileExp           = regexp.MustCompile(`^/(?P<file>.*)`)
)

// PathComponents interface
type PathComponents interface {
	Endpoint() string
	GetEndpointType() EndpointType
	GetName() string
}

type endpointComponents struct {
	endpointType EndpointType
	matchType    matchPattern
	user         string
	repo         string
	blob         string
	branch       string
	path         string
	folder       string
	file         string
	isFolder     bool
}

func patternMatch(endpoint string) *endpointComponents {
	endpointBytes := []byte(endpoint)

	switch {
	case urlExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchURLPattern}).parseFrom(endpoint, urlExp)
	case urlWithoutHostExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchURLWithoutHostPattern}).parseFrom(endpoint, urlWithoutHostExp)
	case fileExp.Match(endpointBytes):
		return (&endpointComponents{matchType: matchFilePattern}).parseFrom(endpoint, fileExp)
	default:
		// endpoint is empty string
		return nil
	}
}

// NewEndpointComponents returns endpointComponents instance
func NewEndpointComponents(path string, referer string) PathComponents {
	ec, code := newEndpointComponentsProxy(path, referer)

	ec.endpointType = code
	return ec
}

func newEndpointComponentsProxy(path string, referer string) (*endpointComponents, EndpointType) {
	refEC := patternMatch(referer)
	pathEC := patternMatch(path)

	switch {
	// path ""
	// 	It is impossible, browser always append "/" to root path
	case pathEC == nil:
		panic("impossible path value")
	// referer ""
	case refEC == nil:
		// path "/user/repo/blob/master/example::/sub/path/"
		// path "/user/repo/blob/master/example::/sub/path"
		// path "/user/repo/blob/master/example/sub/path/"
		// path "/user/repo/blob/master/example/sub/path"
		// path "/example"
		// 	Redirect to index page
		if pathEC.isFolder {
			return pathEC, EndpointRedirect
		}
		switch pathEC.matchType {
		case matchURLPattern, matchURLWithoutHostPattern:
			return pathEC, EndpointRemoteAsset
		default:
			return pathEC, EndpointLocalAsset
		}
	// path is not "", referer is not ""
	default:
		if refEC.isFolder {
			panic("refEC must not be a folder, path will always be redirected to index page")
		}
		switch refEC.matchType<<mpc | pathEC.matchType {

		case matchURLPattern<<mpc | matchURLPattern, matchURLPattern<<mpc | matchURLWithoutHostPattern:
			// referer "/user/repo/blob/master::/index.html"
			// path "./favicon.ico" --> "/user/repo/blob/master::/favicon.ico"
			// path "/user/repo2/blob/master::/favicon.ico"
			// path "/user/repo/blob/master/favicon.ico"
			//	Request remote asset
			return pathEC, EndpointRemoteAsset
		case matchURLPattern<<mpc | matchFilePattern:
			// referer "/user/repo/blob/master::/index.html"
			// path "../favicon.ico" --> "/user/repo/blob/favicon.ico"
			// path "/user/repo/blob/favicon.ico"
			// 	We do not allow path eliminating chunks of github branch pattern
			if refEC.hasSubPrefixWith(path) {
				return pathEC, EndpointNotFound
			}
			// path "/favicon.icon"
			// 	We redirect to remote asset:
			return refEC.setFolder("").setFile(pathEC.getFile()), EndpointRedirect

		case matchURLWithoutHostPattern<<mpc | matchURLPattern, matchURLWithoutHostPattern<<mpc | matchURLWithoutHostPattern:
			// referer "/user/repo/blob/master/index.html"
			// path "/user/repo/blob/master::/index."
			// path "/user/repo/blob/master/favicon.ico"
			return pathEC, EndpointRemoteAsset
		case matchURLWithoutHostPattern<<mpc | matchFilePattern:
			// referer "/user/repo/blob/master/index.html"
			// path "../favicon.ico" --> "/user/repo/blob/master/favicon.ico"
			// 	We do not allow path eliminating chunks of github branch pattern
			if refEC.hasSubPrefixWith(path) {
				return pathEC, EndpointNotFound
			}
			// path "/favicon.ico"
			return refEC.setFile(pathEC.getFile()), EndpointRedirect

		case matchFilePattern<<mpc | matchURLPattern, matchFilePattern<<mpc | matchURLWithoutHostPattern:
			// referer "/index.html"
			// path "/user/repo/blob/master::/index.html"
			// path "/user/repo/blob/master/index.html"
			return pathEC, EndpointRemoteAsset
		case matchFilePattern<<mpc | matchFilePattern:
			// referer "/index.html"
			// path "/"
			// path "/example"
			// 	Redirect to index page
			if pathEC.isFolder {
				return pathEC, EndpointRedirect
			}
			// path "./favicon.icon" -> "/favicon.icon"
			// path "/favicon.icon"
			return pathEC, EndpointLocalAsset
		default:
			panic("unexpected situation")
		}
	}
}

func (ec *endpointComponents) setFile(file string) *endpointComponents {
	ec.file = file
	return ec
}
func (ec *endpointComponents) getFile() string {
	return ec.file
}
func (ec *endpointComponents) setFolder(folder string) *endpointComponents {
	ec.folder = folder
	return ec
}

func (ec *endpointComponents) getFolder() string {
	return ec.folder
}

func (ec *endpointComponents) GetName() string {
	return ec.file
}

func (ec *endpointComponents) parseFrom(endpoint string, reg *regexp.Regexp) *endpointComponents {
	match := reg.FindStringSubmatch(endpoint)

	for i, name := range reg.SubexpNames() {
		switch name {
		case "user":
			ec.user = match[i]
		case "blob":
			ec.blob = match[i]
		case "repo":
			ec.repo = match[i]
		case "branch":
			ec.branch = match[i]
		case "path":
			ec.path = match[i]
			if strings.HasSuffix(ec.path, "/") {
				ec.path = ec.path[0 : len(ec.path)-1]
			}
		case "folder":
			ec.folder = match[i]
			if ec.folder != "" {
				ec.folder = "/" + ec.folder
			}
			if strings.HasSuffix(ec.folder, "/") {
				ec.folder = ec.folder[0 : len(ec.folder)-1]
			}
		case "file":
			ec.file = match[i]
			if path.Ext(ec.file) == "" {
				ec.isFolder = true
				ec.file += "/index.html"
			} else if ec.file == "" {
				ec.file = "/index.html"
			} else if !strings.HasSuffix(ec.file, "/") {
				ec.file = "/" + ec.file
			}
		}
	}

	return ec
}

func (ec *endpointComponents) Endpoint() (endpoint string) {
	switch ec.endpointType {
	case EndpointLocalAsset:
		endpoint = "/" + ec.file
	case EndpointRemoteAsset:
		endpoint = fmt.Sprintf(
			"/%s%s%s%s%s%s",
			ec.user, ec.repo, ec.branch, ec.path, ec.folder, ec.file,
		)
	case EndpointRedirect:
		if ec.matchType == matchURLPattern {
			endpoint = fmt.Sprintf("/%s%s%s%s%s::/%s%s", ec.user, ec.repo, ec.blob, ec.branch, ec.path, ec.folder, ec.file)
		} else {
			endpoint = fmt.Sprintf("/%s%s%s%s%s%s%s", ec.user, ec.repo, ec.blob, ec.branch, ec.path, ec.folder, ec.file)
		}
	}

	return path.Clean(endpoint)
}

func (ec *endpointComponents) GetEndpointType() EndpointType {
	return ec.endpointType
}

// Rule:
// endpoint: /a/b/c/d/, has sub endpoints /a/, /a/b/, /a/b/c/
// testEndpoint: /a'/b'/c'/d'/e', has 6 chunks, returns false
// testEndpoint: /a'/b'/c'/d', has 5 chunks, if /a'/b'/c'/d' contains /a/b/c/ then returns true else returns false
// testEndpoint: /a'/b'/c', has 4 chunks, if /a'/b'/c' contains /a/b/ then returns true else returns false
// testEndpoint: /a'/b', has 3 chunks, if /a'/b' contains /a/ then returns true else returns false
// testEndpoint: /a', has 2 chunks, always returns false
func (ec *endpointComponents) hasSubPrefixWith(endpoint string) bool {
	chunks := strings.Split(endpoint, "/")
	cLen := len(chunks)
	if cLen < 3 {
		return false
	}
	// We only check max four chunks
	if cLen > 5 {
		cLen = 5
	}

	prefix := strings.Join(
		([]string{"/", ec.user, ec.repo, ec.blob})[0:cLen-1],
		"",
	)

	return strings.HasPrefix(endpoint, prefix)
}
