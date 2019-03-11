package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

type structForURLComponentsFunctionality struct {
	testName    string
	testPath    string
	testReferer string
	path        string
	host        string
	err         error
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
		"path is relative path to host; referer is regular",
		"/asset.css",
		"/user/repo/blob/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example",
		nil,
	},
	{
		"no 'blob' in path",
		"/user/repo/master/example//sub/path/index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example",
		nil,
	},
	{
		"no 'blob' in referer path",
		"/asset.css",
		"/user/repo/master/example//sub/path/index.html",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example",
		nil,
	},
	{
		"'tree' in path",
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
		"path is relative path to host; referer without folder",
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
		"path is relative path to host; referer with special char symbols",
		"/asset.css",
		"/user/repo-name/blob/master/example//index.html",
		"/user/repo-name/master/example/asset.css",
		"/user/repo-name/master/example",
		nil,
	},

	{
		"muti chunks path",
		"/user/repo/blob/master/example/sub/path//index.html",
		"",
		"/user/repo/master/example/sub/path/index.html",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		"path is relative path to host; referer with muti chunks path",
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
		"path is relative path to host; referer without folder",
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
		"path is relative path to host; referer without file",
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
		"path is relative path to host; referer without file",
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
		"path is relative path to host; referer without host",
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
		"path is relative path to host; referer without host and path",
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
		"path is relative path to host; referer without host and file",
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
		"path is relative path to host; referer without host and file - 2",
		"/asset.css",
		"/user/repo/blob/master/example/sub/path/",
		"/user/repo/master/example/sub/path/asset.css",
		"/user/repo/master/example/sub/path",
		nil,
	},
	{
		testName: "local file triggers error",
		testPath: "/favicon.ico",
		err:      utils.ErrNotMatchURLPattern,
	},
}

func TestInputCases(t *testing.T) {
	for _, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		if !utils.ExportedBaseExp.Match([]byte(test.testPath)) && !utils.ExportedBaseExp.Match([]byte(test.testReferer)) && test.err == nil {
			t.Errorf("You provided invalid test case in [%s]: got testPath `%s` and testReferer `%s`", descr, test.testPath, test.testReferer)
		}
	}
}

type structForPrevent301RedirectionFunctionality struct {
	testName string
	testPath string
}

var testsForPrevent301RedirectionFunctionality = func() (tests []structForPrevent301RedirectionFunctionality) {
	for _, test := range testsForURLComponentsFunctionality {
		// We chose no referer cases
		if test.testReferer == "" && test.err == nil {
			tests = append(tests, structForPrevent301RedirectionFunctionality{
				testName: test.testName,
				testPath: test.testPath,
			})
		}
	}

	return
}()
