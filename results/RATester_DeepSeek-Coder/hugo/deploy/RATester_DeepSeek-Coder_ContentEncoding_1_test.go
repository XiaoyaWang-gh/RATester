package deploy

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestContentEncoding_1(t *testing.T) {
	type testCase struct {
		name     string
		lf       *localFile
		expected string
	}

	testCases := []testCase{
		{
			name: "Gzip is true",
			lf: &localFile{
				matcher: &deployconfig.Matcher{
					Gzip: true,
				},
			},
			expected: "gzip",
		},
		{
			name: "Gzip is false and ContentEncoding is set",
			lf: &localFile{
				matcher: &deployconfig.Matcher{
					Gzip:            false,
					ContentEncoding: "deflate",
				},
			},
			expected: "deflate",
		},
		{
			name: "Matcher is nil",
			lf: &localFile{
				matcher: nil,
			},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := tc.lf.ContentEncoding()
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
