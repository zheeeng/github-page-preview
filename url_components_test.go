package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		testName string
		testPath string
		path     string
		host     string
	}{
		{
			"regular",
			"/zheeeng/exportfromjson/blob/master/example//index.html",
			"/zheeeng/exportfromjson/master/example/index.html",
			"/zheeeng/exportfromjson/master/example",
		},
		{
			"with special char symbols",
			"/zheeeng/export-from-json/blob/master/example//index.html",
			"/zheeeng/export-from-json/master/example/index.html",
			"/zheeeng/export-from-json/master/example",
		},
		{
			"with special char symbols2",
			"/zheeeng/@roundation/blob/master/example//index.html",
			"/zheeeng/@roundation/master/example/index.html",
			"/zheeeng/@roundation/master/example",
		},
		{
			"long path",
			"/zheeeng/test/blob/master/example/sub/path//index.html",
			"/zheeeng/test/master/example/sub/path/index.html",
			"/zheeeng/test/master/example/sub/path",
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("\nTest %s failed:\n", test.testName)

		pc := pathComponents{}

		pc.parseFrom(test.testPath)

		rawPath := pc.compileToRaw()
		if rawPath != test.testPath {
			t.Errorf("%s[compile to raw]: got %s, want %s", descr, rawPath, test.testPath)
		}

		path := pc.compileToRequestPath()
		if path != test.path {
			t.Errorf("%s[compile to request path]: got %s, want %s", descr, path, test.path)
		}

		host := pc.compileToStaticHost()
		if host != test.host {
			t.Errorf("%s[compile to request host]: got %s, want %s", descr, host, test.host)
		}
	}
}
