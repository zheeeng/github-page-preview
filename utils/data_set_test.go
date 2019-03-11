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
		"/asset.css",
		"/user/repo/blob/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/asset.css",
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
		"/asset.css",
		"/user/repo/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/asset.css",
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
		"/asset.css",
		"/user/repo/tree/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/asset.css",
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
		"/asset.css",
		"/user/repo/blob/master/example//index.html",
		"/user/repo/master/example/asset.css",
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
		"/asset.css",
		"/user/repo-name/blob/master/example//index.html",
		"/user/repo-name/master/example/asset.css",
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
		"/asset.css",
		"/user/repo/blob/master/example/sub/path//index.html",
		"/user/repo/master/example/sub/path/asset.css",
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
		"/asset.css",
		"/user/repo/blob/master//",
		"/user/repo/master/asset.css",
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
		"/asset.css",
		"/user/repo/blob/master/example//sub/path",
		"/user/repo/master/example/sub/path/asset.css",
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
		"/asset.css",
		"/user/repo/blob/master/example//sub/path/",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example",
		nil,
	},

	// specifications fro no hosts URL
	{
		"no host",
		"/user/repo/blob/master/example/sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer without host",
		"/asset.css",
		"/user/repo/blob/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example/sub/path",
		nil,
	},

	{
		"no host, no path",
		"/user/repo/blob/master/index.html",
		"",
		"/user/repo/master/index.html",
		"/user/repo/master",
		nil,
	},
	{
		"endpoint is relative to host; referer without host and path",
		"/asset.css",
		"/user/repo/blob/master/index.html",
		"/user/repo/master/asset.css",
		"/user/repo/master",
		nil,
	},

	{
		"no host, no file",
		"/user/repo/blob/master/example/sub/path",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer without host and file",
		"/asset.css",
		"/user/repo/blob/master/example/sub/path",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example/sub/path",
		nil,
	},

	{
		"no host, no file - 2",
		"/user/repo/blob/master/example/sub/path/",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"endpoint is relative to host; referer without host and file - 2",
		"/asset.css",
		"/user/repo/blob/master/example/sub/path/",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		testName:     "local file triggers error",
		testEndpoint: "/favicon.ico",
		err:          utils.ErrNotMatchURLPattern,
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
