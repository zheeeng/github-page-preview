package utils

import (
	"testing"
)

func Test_hasSubPrefixWith(t *testing.T) {
	tests := []struct {
		testEndpoint string
		testPath     string
		has          bool
	}{
		{
			"/user/repo/blob/master/index.html",
			"/user/repo/blob/master/index.html",
			true,
		},
		{
			"/user/repo/blob/master/example/index.html",
			"/user/repo/blob/master/index.html",
			true,
		},
		{
			"/user/repo/blob/master/index.html",
			"/",
			false,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user",
			false,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo",
			true,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo/",
			true,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo/blob/index.html",
			true,
		},
		{
			"/user/repo/blob/master/example/index.html",
			"/user/repo/blob/index.html",
			true,
		},
		{
			"/user/repo/blob/master/example::/index.html",
			"/user/repo/blob/index.html",
			true,
		},
		{
			"/user/repo/blob/master/example::/sub/index.html",
			"/user/repo/blob/index.html",
			true,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo/tree/index.html",
			false,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo2/blob/index.html",
			false,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user2/repo/blob/index.html",
			false,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo/index.html",
			true,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/repo2/index.html",
			false,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user/index.html",
			true,
		},
		{
			"/user/repo/blob/master/index.html",
			"/user2/index.html",
			false,
		},
	}

	for i, test := range tests {
		ec := patternMatch(test.testEndpoint)
		if has := ec.hasSubPrefixWith(test.testPath); has != test.has {
			hasStr := "has"
			if !test.has {
				hasStr = hasStr + " no"
			}

			t.Errorf(
				"\n[%d] expected `%s` %s sub prefix with `%s`",
				i, test.testEndpoint, hasStr, test.testPath,
			)
		}
	}
}
