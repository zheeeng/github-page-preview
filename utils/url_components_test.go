package utils_test

import (
	"fmt"
	"testing"

	"github.com/github-page-preview/utils"
)

func codeTransToLiteral(code utils.EndpointType) string {
	switch code {
	case EndpointLocalAsset:
		return "EndpointLocalAsset"
	case EndpointRemoteAsset:
		return "EndpointRemoteAsset"
	case EndpointRedirect:
		return "EndpointRedirect"
	case EndpointNotFound:
		return "EndpointNotFound"
	default:
		panic("Debug plz")
	}
}

func TestParsePanic(t *testing.T) {
	for i, test := range testsForURLComponentsPanic {
		func(i int, test structForURLComponentsPanic) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf(
						"\n[%d] [Trigger Panic]: \ntestEndpoint is `%s`, \ntestReferer is `%s`\nshould trigger panic",
						i, test.testEndpoint, test.testReferer,
					)
				}
			}()
			utils.NewEndpointComponents(test.testEndpoint, test.testReferer)
		}(i, test)
	}
}

func TestParse(t *testing.T) {
	var panicLogIndex int
	var panicLogDescr string
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("\n[%d] Test has panic: %s\tPanic is: `%v` ", panicLogIndex, panicLogDescr, r)
		}
	}()

	for i, test := range testsForURLComponentsFunctionality {
		descr := fmt.Sprintf("\n[%d] Test [%s] failed:\n", i, test.testName)
		panicLogIndex = i
		panicLogDescr = descr

		ec := utils.NewEndpointComponents(test.testEndpoint, test.testReferer)

		code := ec.GetEndpointType()

		if code != test.code {
			t.Errorf(
				"%s[Trigger Error]: \ntestEndpoint is `%s`, \ntestReferer is `%s`\ngot `%v`, want `%v`",
				descr, test.testEndpoint, test.testReferer,
				codeTransToLiteral(code), codeTransToLiteral(test.code),
			)
		}

		if test.code == utils.EndpointNotFound {
			continue
		}

		endpoint := ec.Endpoint()

		if endpoint != test.endpointOrRaw {
			t.Errorf(
				"%s[Endpoint]: testEndpoint is `%s`, testReferer is `%s`\ngot `%s`, want `%s`",
				descr, test.testEndpoint, test.testReferer, endpoint, test.endpointOrRaw,
			)
		}
	}
}
