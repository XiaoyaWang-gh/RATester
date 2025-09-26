package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestFileFromFS_1(t *testing.T) {
	testCases := []struct {
		name     string
		filepath string
		fs       http.FileSystem
		expected string
	}{
		{
			name:     "Test case 1",
			filepath: "/test/file",
			fs:       http.Dir("./testdata"),
			expected: "testdata/file",
		},
		{
			name:     "Test case 2",
			filepath: "/test/file2",
			fs:       http.Dir("./testdata"),
			expected: "testdata/file2",
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			// Create a new context
			ctx := &Context{
				Request: &http.Request{
					URL: &url.URL{
						Path: tc.filepath,
					},
				},
			}

			// Call the method under test
			ctx.FileFromFS(tc.filepath, tc.fs)

			// Check that the request path was updated as expected
			if ctx.Request.URL.Path != tc.expected {
				t.Errorf("Expected request path to be %q, but got %q", tc.expected, ctx.Request.URL.Path)
			}
		})
	}
}
