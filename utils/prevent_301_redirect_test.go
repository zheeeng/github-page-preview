package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestConvertions(t *testing.T) {
	var tests = []struct {
		testName string
		testPath string
	}{
		{
			"regular",
			"/user/repo/blob/master/example//sub/path/index.html",
		},
		{
			"no 'blob' in path",
			"/user/repo/master/example//sub/path/index.html",
		},
		{
			"'tree' in path",
			"/user/repo/tree/master/example//sub/path/index.html",
		},
		{
			"no folder",
			"/user/repo/blob/master/example//index.html",
		},
		{
			"with special char symbols",
			"/user/repo-name/blob/master/example//index.html",
		},
		{
			"muti chunks path",
			"/user/repo/blob/master/example/sub/path//index.html",
		},
		{
			"no folder",
			"/user/repo/blob/master//",
		},
		{
			"no file",
			"/user/repo/blob/master/example//sub/path",
		},
		{
			"no file2",
			"/user/repo/blob/master/example//sub/path/",
		},
		// specifications fro no hosts URL
		{
			"no host",
			"/user/repo/blob/master/example/sub/path/index.html",
		},
		{
			"no host, no path",
			"/user/repo/blob/master/index.html",
		},
		{
			"no host, no file",
			"/user/repo/blob/master/example/sub/path",
		},
		{
			"no host, no file - 2",
			"/user/repo/blob/master/example/sub/path/",
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		r := http.Request{
			URL: &url.URL{
				Path: test.testPath,
			},
		}

		PreventRedirection(&r)

		path := RestoreHijacked(r.URL.Path)

		if path != test.testPath {
			t.Errorf("%s[Convert and Restore]: got %v, want %v", descr, path, test.testPath)
		}
	}
}
