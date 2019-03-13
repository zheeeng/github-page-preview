package utils_test

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"testing"

	"github.com/github-page-preview/utils"
)

func TestConvertions(t *testing.T) {
	for _, test := range testsForPrevent301RedirectionFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		r := http.Request{
			URL: &url.URL{
				Path: test.testEndpoint,
			},
		}

		utils.PreventRedirection(&r)

		h := r.URL.Path

		testEndpointMatched := utils.ExportedBaseExp.Match([]byte(test.testEndpoint))

		if matched := utils.ExportedBaseExp.Match([]byte(h)); matched != testEndpointMatched {
			t.Errorf("%s[Convertion breaks URI structure]: got `%s`, should match pattern `%v`", descr, h, utils.ExportedBaseExp.String())
		}

		if matched, _ := regexp.Match(utils.ExportedIndexPattern, []byte(h)); matched {
			t.Errorf("%s[Strip index.html]: got `%s`", descr, h)
		}

		if matched, _ := regexp.Match(utils.ExportedDelimiterPattern, []byte(h)); matched {
			t.Errorf("%s[Strip delimiter]: got `%s`", descr, h)
		}

		endpoint := utils.RestoreHijacked(r.URL.Path)

		if endpoint != test.testEndpoint {
			t.Errorf("%s[Convert and restore]: got `%v`, want `%v`", descr, endpoint, test.testEndpoint)
		}
	}
}
