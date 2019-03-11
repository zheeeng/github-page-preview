package utils_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/github-page-preview/utils"
)

func TestConvertions(t *testing.T) {
	for _, test := range testsForPrevent301RedirectionFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		r := http.Request{
			URL: &url.URL{
				Path: test.testPath,
			},
		}

		utils.PreventRedirection(&r)

		path := utils.RestoreHijacked(r.URL.Path)

		if path != test.testPath {
			t.Errorf("%s[Convert and Restore]: got %v, want %v", descr, path, test.testPath)
		}
	}
}
