package crd

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsWatchedNamespace_3(t *testing.T) {
	testCases := []struct {
		desc     string
		isAll    bool
		ns       string
		expected bool
	}{
		{
			desc:     "is all",
			isAll:    true,
			ns:       "test",
			expected: true,
		},
		{
			desc:     "is not all",
			isAll:    false,
			ns:       "test",
			expected: false,
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

			c := &clientWrapper{
				isNamespaceAll: test.isAll,
				watchedNamespaces: []string{
					"test",
				},
			}

			actual := c.isWatchedNamespace(test.ns)
			assert.Equal(t, test.expected, actual)
		})
	}
}
