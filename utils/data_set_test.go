package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

type structForURLComponentsFunctionality struct {
	testName     string
	testEndpoint string
	testReferer  string
	endpoint     string
	host         string
	err          error
}

var testsForURLComponentsFunctionality = []structForURLComponentsFunctionality{
	{
		"regular",
		"/user/repo/blob/master/example//sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"endpoint is relative to host; referer is regular",
		"/favicon.ico",
		"/user/repo/blob/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example",
		nil,
	},
	{
		"no 'blob/tree' in endpoint",
		"/user/repo/master/example//sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"no 'blob/tree' in referer",
		"/favicon.ico",
		"/user/repo/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example",
		nil,
	},
	{
		"'tree' in endpoint",
		"/user/repo/tree/master/example//sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"'tree' in referer",
		"/favicon.ico",
		"/user/repo/tree/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example",
		nil,
	},
	{
		"no folder",
		"/user/repo/blob/master/example//index.html",
		"",
		"/user/repo/master/example/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"endpoint is relative to host; referer without folder",
		"/favicon.ico",
		"/user/repo/blob/master/example//index.html",
		"/user/repo/master/example/favicon.ico",
		"/user/repo/master/example",
		nil,
	},

	{
		"with special char symbols",
		"/user/repo-name/blob/master/example//index.html",
		"",
		"/user/repo-name/master/example/index.html",
		"/user/repo-name/master/example",
		nil,
	},
	{
		"endpoint is relative to host; referer with special char symbols",
		"/favicon.ico",
		"/user/repo-name/blob/master/example//index.html",
		"/user/repo-name/master/example/favicon.ico",
		"/user/repo-name/master/example",
		nil,
	},

	{
		"muti chunks",
		"/user/repo/blob/master/example/sub/path//index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer with muti chunks",
		"/favicon.ico",
		"/user/repo/blob/master/example/sub/path//index.html",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example/sub/path",
		nil,
	},

	{
		"no folder",
		"/user/repo/blob/master//",
		"",
		"/user/repo/master/index.html",
		"/user/repo/master",
		nil,
	},
	{
		"endpoint is relative to host; referer without folder",
		"/favicon.ico",
		"/user/repo/blob/master//",
		"/user/repo/master/favicon.ico",
		"/user/repo/master",
		nil,
	},

	{
		"no file",
		"/user/repo/blob/master/example//sub/path",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"endpoint is relative to host; referer without file",
		"/favicon.ico",
		"/user/repo/blob/master/example//sub/path",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example",
		nil,
	},
	{
		"no file2",
		"/user/repo/blob/master/example//sub/path/",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"endpoint is relative to host; referer without file",
		"/favicon.ico",
		"/user/repo/blob/master/example//sub/path/",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example",
		nil,
	},

	// specifications fro no specified hosts URL
	{
		"no specified host",
		"/user/repo/blob/master/example/sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer without specified host",
		"/favicon.ico",
		"/user/repo/blob/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example/sub/path",
		nil,
	},

	{
		"no specified host, no path",
		"/user/repo/blob/master/index.html",
		"",
		"/user/repo/master/index.html",
		"/user/repo/master",
		nil,
	},
	{
		"endpoint is relative to host; referer without host and path",
		"/favicon.ico",
		"/user/repo/blob/master/index.html",
		"/user/repo/master/favicon.ico",
		"/user/repo/master",
		nil,
	},

	{
		"no specified host, no file",
		"/user/repo/blob/master/example/sub/path",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer without specified host and file",
		"/favicon.ico",
		"/user/repo/blob/master/example/sub/path",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example/sub/path",
		nil,
	},

	{
		"no specified host, no file - 2",
		"/user/repo/blob/master/example/sub/path/",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer without specified host and file - 2",
		"/favicon.ico",
		"/user/repo/blob/master/example/sub/path/",
		"/user/repo/master/example/sub/path/favicon.ico",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		testName:     "local root page triggers error",
		testEndpoint: "/",
		err:          utils.ErrNotRecognizeURL,
	},
	{
		testName:     "local index page triggers error",
		testEndpoint: "/index.html",
		err:          utils.ErrNotRecognizeURL,
	},
	{
		testName:     "local asset, referer is root page",
		testEndpoint: "/favicon.ico",
		testReferer:  "/",
		err:          utils.ErrNotRecognizeURL,
	},
	{
		testName:     "local asset, referer is index page",
		testEndpoint: "/favicon.ico",
		testReferer:  "/index.html",
		err:          utils.ErrNotRecognizeURL,
	},
}

func TestInputCases(t *testing.T) {
	for _, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		if !utils.ExportedBaseExp.Match([]byte(test.testEndpoint)) && !utils.ExportedBaseExp.Match([]byte(test.testReferer)) && test.err == nil {
			t.Errorf("You provided invalid test case in [%s]: got testEndpoint `%s` and testReferer `%s`", descr, test.testEndpoint, test.testReferer)
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
		if test.testReferer == "" && test.err == nil {
			tests = append(tests, structForPrevent301RedirectionFunctionality{
				testName:     test.testName,
				testEndpoint: test.testEndpoint,
			})
		}
	}

	return
}()
