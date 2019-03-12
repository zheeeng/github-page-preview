package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

func TestParse(t *testing.T) {
	for _, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\nTest [%s] failed:\n", test.testName)

		ec, err := utils.NewEndpointComponents(test.testEndpoint, test.testReferer)

		if err != test.err {
			t.Errorf(
				"%s[ErrorTrigger]: testEndpoint is `%s`, testReferer is `%s`\ngot `%v`, want `%v`",
				descr, test.testEndpoint, test.testReferer, err, test.err,
			)
		}

		if err != nil {
			continue
		}

		endpoint := ec.Endpoint()
		if endpoint != test.endpoint {
			t.Errorf(
				"%s[Endpoint]: testEndpoint is `%s`, testReferer is `%s`\ngot `%s`, want `%s`",
				descr, test.testEndpoint, test.testReferer, endpoint, test.endpoint,
			)
		}

		host := ec.StaticHost()
		if host != test.host {
			t.Errorf(
				"%s[StaticHost]: testEndpoint is `%s`, testReferer is `%s`\ngot `%s`, want `%s`",
				descr, test.testEndpoint, test.testReferer, host, test.host,
			)
		}
	}
}
