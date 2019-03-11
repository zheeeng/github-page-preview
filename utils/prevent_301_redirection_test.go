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
				Path: test.testPath,
			},
		}

		utils.PreventRedirection(&r)

		h := r.URL.Path

		if matched := utils.ExportedBaseExp.Match([]byte(h)); !matched {
			t.Errorf("%s[Convertion breaks URI structure]: got %s, should match pattern `%v`", descr, h, utils.ExportedBaseExp.String())
		}

		if matched, _ := regexp.Match(utils.ExportedIndexPattern, []byte(h)); matched {
			t.Errorf("%s[Strip index.html]: got %s", descr, h)
		}

		if matched, _ := regexp.Match(utils.ExportedSuffixSlashPattern, []byte(h)); matched {
			t.Errorf("%s[Strip suffix slash]: got %s", descr, h)
		}

		if matched, _ := regexp.Match(utils.ExportedDelimiterPattern, []byte(h)); matched {
			t.Errorf("%s[Strip double-slash delimiter]: got %s", descr, h)
		}

		path := utils.RestoreHijacked(r.URL.Path)

		if path != test.testPath {
			t.Errorf("%s[Convert and restore]: got %v, want %v", descr, path, test.testPath)
		}
	}
}