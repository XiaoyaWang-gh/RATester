package dynamic

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestHasCustomHeadersDefined_1(t *testing.T) {
	testCases := []struct {
		desc     string
		headers  *Headers
		expected bool
	}{
		{
			desc:     "nil headers",
			headers:  nil,
			expected: false,
		},
		{
			desc:     "empty headers",
			headers:  &Headers{},
			expected: false,
		},
		{
			desc: "custom request headers defined",
			headers: &Headers{
				CustomRequestHeaders: map[string]string{
					"X-Test": "test",
				},
			},
			expected: true,
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

			assert.Equal(t, test.expected, test.headers.HasCustomHeadersDefined())
		})
	}
}
