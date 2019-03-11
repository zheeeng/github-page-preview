package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

func TestParse(t *testing.T) {
	for _, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		pc, err := utils.NewPathComponents(test.testEndpoint, test.testReferer)

		if err != test.err {
			t.Errorf("%s[ErrorTrigger]: got %v, want %v", descr, err, test.err)
		}

		if err != nil {
			continue
		}

		endpoint := pc.Endpoint()
		if endpoint != test.endpoint {
			t.Errorf("%s[Endpoint]: got %s, want %s", descr, endpoint, test.endpoint)
		}

		host := pc.StaticHost()
		if host != test.host {
			t.Errorf("%s[StaticHost]: got %s, want %s", descr, host, test.host)
		}
	}
}
