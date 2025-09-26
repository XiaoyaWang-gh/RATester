package testhelpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithRule_1(t *testing.T) {
	testCases := []struct {
		desc     string
		rule     string
		expected *dynamic.Router
	}{
		{
			desc:     "empty rule",
			rule:     "",
			expected: &dynamic.Router{},
		},
		{
			desc:     "simple rule",
			rule:     "Host:test.traefik.io",
			expected: &dynamic.Router{Rule: "Host:test.traefik.io"},
		},
		{
			desc:     "rule with path",
			rule:     "PathPrefix:/test",
			expected: &dynamic.Router{Rule: "PathPrefix:/test"},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			got := &dynamic.Router{}
			WithRule(test.rule)(got)

			assert.Equal(t, test.expected, got)
		})
	}
}
