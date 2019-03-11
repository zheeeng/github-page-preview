package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

func TestParse(t *testing.T) {
	for _, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		pc, err := utils.NewPathComponents(test.testPath, test.testReferer)

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
