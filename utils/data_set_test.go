package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

const EndpointLocalAsset = utils.EndpointLocalAsset
const EndpointRemoteAsset = utils.EndpointRemoteAsset
const EndpointRedirect = utils.EndpointRedirect
const EndpointNotFound = utils.EndpointNotFound

type structForURLComponentsPanic struct {
	testEndpoint string
	testReferer  string
}

var testsForURLComponentsPanic = []structForURLComponentsPanic{
	{"", "/user/repo/blob/master/example/sub/path/index.html"},
	{"/", "/"},
	{"/index.html", "/"},
	{"/user/repo/blob/master/example::/sub/path/index.html", "/user/repo/blob/master/example/sub/path/"},
	{"/user/repo/blob/master/example::/sub/path/index.html", "/user/repo/blob/master/example/sub/path"},
	{"/user/repo/blob/master/example/sub/path/index.html", "/user/repo/blob/master/example/sub/path/"},
	{"/user/repo/blob/master/example/sub/path/index.html", "/user/repo/blob/master/example/sub/path"},
}

type structForURLComponentsFunctionality struct {
	testName      string
	testEndpoint  string
	testReferer   string
	endpointOrRaw string
	code          utils.EndpointType
}

var testsForURLComponentsFunctionality = []structForURLComponentsFunctionality{
	{
		"regular",
		"/user/repo/blob/master/example::/sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		EndpointRemoteAsset,
	},
	{
		"endpoint redirects to referer; referer is regular",
		"/",
		"/user/repo/blob/master/example::/sub/path/index.html",
		"/user/repo/blob/master/example::/index.html",
		EndpointRedirect,
	},
	{
		"endpoint is relative to host; referer is regular",
		"/favicon.ico",
		"/user/repo/blob/master/example::/sub/path/index.html",
		"/user/repo/blob/master/example::/favicon.ico",
		EndpointRedirect,
	},
	{
		"endpoint backs to the ancestors of referer; referer is regular",
		// "/user/repo/blob/master/example::/sub/path/index.html" + "../bundle.js"
		"/user/repo/blob/master/example::/sub/bundle.js",
		"/user/repo/blob/master/example::/sub/path/index.html",
		"/user/repo/master/example/sub/bundle.js",
		EndpointRemoteAsset,
	},
	{
		"endpoint backs to the ancestors of referer; referer is regular - 2",
		// "/user/repo/blob/master/example::/sub/path/index.html" + "../../bundle.js"
		"/user/repo/blob/master/example::/bundle.js",
		"/user/repo/blob/master/example::/sub/path/index.html",
		"/user/repo/master/example/bundle.js",
		EndpointRemoteAsset,
	},
	{
		"endpoint backs out of the referer's host; referer is regular",
		// "/user/repo/blob/master/example::/sub/path/index.html" + "../../../bundle.js",
		"/user/repo/blob/master/bundle.js",
		"/user/repo/blob/master/example::/sub/path/index.html",
		"/user/repo/master/bundle.js",
		EndpointRemoteAsset,
	},
	{
		"'tree' in endpoint",
		"/user/repo/tree/master/example::/sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		EndpointRemoteAsset,
	},
	{
		"no folder",
		"/user/repo/blob/master/example::/index.html",
		"",
		"/user/repo/master/example/index.html",
		EndpointRemoteAsset,
	},
	{
		"endpoint is relative to host; referer has no folder",
		"/favicon.ico",
		"/user/repo/blob/master/example::/index.html",
		"/user/repo/blob/master/example::/favicon.ico",
		EndpointRedirect,
	},

	{
		"with special char symbols",
		"/user/repo-name/blob/master/example::/index.html",
		"",
		"/user/repo-name/master/example/index.html",
		EndpointRemoteAsset,
	},
	{
		"endpoint is relative to host; referer with special char symbols",
		"/favicon.ico",
		"/user/repo-name/blob/master/example::/index.html",
		"/user/repo-name/blob/master/example::/favicon.ico",
		EndpointRedirect,
	},

	{
		"muti chunks",
		"/user/repo/blob/master/example/sub/path::/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		EndpointRemoteAsset,
	},
	{
		"endpoint is relative to host; referer with muti chunks",
		"/favicon.ico",
		"/user/repo/blob/master/example/sub/path::/index.html",
		"/user/repo/blob/master/example/sub/path::/favicon.ico",
		EndpointRedirect,
	},

	{
		"no folder, no file",
		"/user/repo/blob/master::/",
		"",
		"/user/repo/blob/master::/index.html",
		EndpointRedirect,
	},

	{
		"no file",
		"/user/repo/blob/master/example::/sub/path",
		"",
		"/user/repo/blob/master/example::/sub/path/index.html",
		EndpointRedirect,
	},
	{
		"no file - 2",
		"/user/repo/blob/master/example::/sub/path/",
		"",
		"/user/repo/blob/master/example::/sub/path/index.html",
		EndpointRedirect,
	},

	// specifications for no specified hosts URL
	{
		"no specified host",
		"/user/repo/blob/master/example/sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		EndpointRemoteAsset,
	},
	{
		"endpoint redirects to referer; referer has no specified host",
		"/",
		"/user/repo/blob/master/example/sub/path/index.html",
		"/user/repo/blob/master/example/sub/path/index.html",
		EndpointRedirect,
	},
	{
		"endpoint is relative to host; referer has no specified host",
		"/favicon.ico",
		"/user/repo/blob/master/example/sub/path/index.html",
		"/user/repo/blob/master/example/sub/path/favicon.ico",
		EndpointRedirect,
	},
	{
		"endpoint backs to the ancestors of referer; referer has no specified host",
		// "/user/repo/blob/master/example/sub/path/index.html" + "../../bundle.js"
		"/user/repo/blob/master/example/bundle.js",
		"/user/repo/blob/master/example/sub/path/index.html",
		"/user/repo/master/example/bundle.js",
		EndpointRemoteAsset,
	},
	{
		"endpoint backs to the ancestors of referer; referer has no specified host - 2",
		// "/user/repo/blob/master/example/sub/path/index.html" + "../../../bundle.js",
		"/user/repo/blob/master/bundle.js",
		"/user/repo/blob/master/example/sub/path/index.html",
		"/user/repo/master/bundle.js",
		EndpointRemoteAsset,
	},
	{
		testName: "endpoint backs to branch chunk; referer has no specified host",
		// "/user/repo/blob/master/example/sub/path/index.html" + "../../../../bundle.js",
		testEndpoint: "/user/repo/blob/bundle.js",
		testReferer:  "/user/repo/blob/master/example/sub/path/index.html",
		code:         EndpointNotFound,
	},

	{
		"no specified host, no path",
		"/user/repo/blob/master/index.html",
		"",
		"/user/repo/master/index.html",
		EndpointRemoteAsset,
	},
	{
		"endpoint is relative to host; referer has no specified host and path",
		"/favicon.ico",
		"/user/repo/blob/master/index.html",
		"/user/repo/blob/master/favicon.ico",
		EndpointRedirect,
	},
	{
		"endpoint backs to the ancestors of referer; referer has no specified host and path",
		"/user/repo/blob/master/bundle.js",
		"/user/repo/blob/master/index.html",
		"/user/repo/master/bundle.js",
		EndpointRemoteAsset,
	},
	{
		testName: "endpoint backs to branch chunk; referer has no specified host and path",
		// "/user/repo/blob/master/index.html" + "../bundle.js",
		testEndpoint: "/user/repo/blob/bundle.js",
		testReferer:  "/user/repo/blob/master/index.html",
		code:         EndpointNotFound,
	},

	{
		"no specified host, no file - 2",
		"/user/repo/blob/master/example/sub/path",
		"",
		"/user/repo/blob/master/example/sub/path/index.html",
		EndpointRedirect,
	},
	{
		"no specified host, no file - 2",
		"/user/repo/blob/master/example/sub/path/",
		"",
		"/user/repo/blob/master/example/sub/path/index.html",
		EndpointRedirect,
	},
	{
		"endpoint redirects to referer; referer is index page",
		"/",
		"/index.html",
		"/index.html",
		EndpointRedirect,
	},
	{
		"local asset; referer is root page",
		"/favicon.ico",
		"/index.html",
		"/favicon.ico",
		EndpointLocalAsset,
	},
	{
		"local asset, endpoint is folder",
		"/example",
		"",
		"/example/index.html",
		EndpointRedirect,
	},
	{
		"local asset, endpoint is folder - 2",
		"/example",
		"",
		"/example/index.html",
		EndpointRedirect,
	},
	{
		"local asset, endpoint is folder; referer is index page",
		"/example",
		"/index.html",
		"/example/index.html",
		EndpointRedirect,
	},
}

func TestInputCases(t *testing.T) {
	for _, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		if (test.code == utils.EndpointRemoteAsset || test.code == utils.EndpointRedirect) &&
			!utils.ExportedBaseExp.Match([]byte(test.testEndpoint)) &&
			!utils.ExportedBaseExp.Match([]byte(test.testReferer)) {
			t.Errorf(
				"%sYou provided invalid test case: testEndpoint `%s` and testReferer `%s`",
				descr, test.testEndpoint, test.testReferer,
			)
		}
	}
}

type structForPrevent301RedirectionFunctionality struct {
	testName     string
	testEndpoint string
}

var testsForPrevent301RedirectionFunctionality = func() (tests []structForPrevent301RedirectionFunctionality) {
	for _, test := range testsForURLComponentsFunctionality {
		// We chose no referer cases
		if test.testReferer == "" {
			tests = append(tests, structForPrevent301RedirectionFunctionality{
				testName:     test.testName,
				testEndpoint: test.testEndpoint,
			})
		}
	}

	return
}()
