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
	}{
		{
			"regular",
			"/zheeeng/exportfromjson/blob/master/example//index.html",
			"/zheeeng/exportfromjson/master/example/index.html",
		},
		{
			"with special char symbols",
			"/zheeeng/export-from-json/blob/master/example//index.html",
			"/zheeeng/export-from-json/master/example/index.html",
		},
		{
			"with special char symbols2",
			"/zheeeng/@roundation/blob/master/example//index.html",
			"/zheeeng/@roundation/master/example/index.html",
		},
		{
			"long path",
			"/zheeeng/test/blob/master/example/sub/path//index.html",
			"/zheeeng/test/master/example/sub/path/index.html",
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("\nTest %s failed:\n", test.testName)

		pc := pathComponents{}

		pc.parseFrom(test.testPath)

		rawPath := pc.compileToRawPath()
		if rawPath != test.testPath {
			t.Errorf("%s[compile to raw path]: got %s, want %s", descr, rawPath, test.testPath)
		}

		path := pc.compileToPath()
		if path != test.path {
			t.Errorf("%s[compile to path]: got %s, want %s", descr, path, test.path)
		}
	}
}
