package main

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
	}{
		{
			"regular",
			"/user/repo/blob/master/example//index.html",
			"",
			"/user/repo/master/example/index.html",
			"/user/repo/master/example",
		},
		{
			"with special char symbols",
			"/user/repo-name/blob/master/example//index.html",
			"",
			"/user/repo-name/master/example/index.html",
			"/user/repo-name/master/example",
		},
		{
			"muti chunks path",
			"/user/repo/blob/master/example/sub/path//index.html",
			"",
			"/user/repo/master/example/sub/path/index.html",
			"/user/repo/master/example/sub/path",
		},
		{
			"no path",
			"/user/repo/blob/master//",
			"",
			"/user/repo/master/index.html",
			"/user/repo/master",
		},
		{
			"no asset",
			"/user/repo/blob/master/example/sub/path//",
			"",
			"/user/repo/master/example/sub/path/index.html",
			"/user/repo/master/example/sub/path",
		},
		{
			"no file",
			"/user/repo/blob/master//example",
			"",
			"/user/repo/master/example/index.html",
			"/user/repo/master",
		},
		{
			"no file - 2",
			"/user/repo/blob/master//example/",
			"",
			"/user/repo/master/example/index.html",
			"/user/repo/master",
		},
		{
			"multi chunks asset",
			"/user/repo/blob/master/example//sub/path/index.html",
			"",
			"/user/repo/master/example/sub/path/index.html",
			"/user/repo/master/example",
		},
		{
			"no host",
			"/user/repo/blob/master/example/sub/path/index.html",
			"",
			"/user/repo/master/example/sub/path/index.html",
			"/user/repo/master/example/sub/path",
		},
		{
			"no host, no path",
			"/user/repo/blob/master/index.html",
			"",
			"/user/repo/master/index.html",
			"/user/repo/master",
		},
		{
			"no host, no file",
			"/user/repo/blob/master/example/sub/path",
			"",
			"/user/repo/master/example/sub/path/index.html",
			"/user/repo/master/example/sub/path",
		},
		{
			"no host, no file - 2",
			"/user/repo/blob/master/example/sub/path/",
			"",
			"/user/repo/master/example/sub/path/index.html",
			"/user/repo/master/example/sub/path",
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		pc := NewPathComponents(test.testPath, test.testReferer)

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