package dynamic

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestHasCorsHeadersDefined_1(t *testing.T) {
	testCases := []struct {
		name     string
		headers  *Headers
		expected bool
	}{
		{
			name: "nil",
		},
		{
			name: "empty",
			headers: &Headers{
				AccessControlAllowCredentials: false,
				AccessControlAllowHeaders:     []string{},
				AccessControlAllowMethods:     []string{},
				AccessControlAllowOriginList:  []string{},
				AccessControlExposeHeaders:    []string{},
			},
		},
		{
			name: "defined",
			headers: &Headers{
				AccessControlAllowCredentials: true,
				AccessControlAllowHeaders:     []string{"foo"},
				AccessControlAllowMethods:     []string{"bar"},
				AccessControlAllowOriginList:  []string{"baz"},
				AccessControlExposeHeaders:    []string{"qux"},
			},
			expected: true,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			assert.Equal(t, test.expected, test.headers.HasCorsHeadersDefined())
		})
	}
}
