package utils

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		testName    string
		testPath    string
		testReferer string
		path        string
		host        string
		err         error
	}{
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
			err:      ErrNotRecognize,
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		pc, err := NewPathComponents(test.testPath, test.testReferer)

		if err != test.err {
			t.Errorf("%s[ErrorTrigger]: got %v, want %v", descr, err, test.err)
		}

		if err != nil {
			continue
		}

		path := pc.RequestPath()
		if path != test.path {
			t.Errorf("%s[RequestPath]: got %s, want %s", descr, path, test.path)
		}

		host := pc.StaticHost()
		if host != test.host {
			t.Errorf("%s[StaticHost]: got %s, want %s", descr, host, test.host)
		}
	}
}
