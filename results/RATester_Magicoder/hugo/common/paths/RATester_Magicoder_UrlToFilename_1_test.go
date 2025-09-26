package paths

import (
	"fmt"
	"testing"
)

func TestUrlToFilename_1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		ok       bool
	}{
		{
			name:     "valid url",
			input:    "http://example.com/path/to/file",
			expected: "HTTP:EXAMPLE.COM/PATH/TO/FILE",
			ok:       true,
		},
		{
			name:     "invalid url",
			input:    "invalid url",
			expected: "invalid url",
			ok:       false,
		},
		{
			name:     "empty url",
			input:    "",
			expected: "",
			ok:       false,
		},
		{
			name:     "url with query",
			input:    "http://example.com/path/to/file?query=param",
			expected: "HTTP:EXAMPLE.COM/PATH/TO/FILE",
			ok:       true,
		},
		{
			name:     "url with fragment",
			input:    "http://example.com/path/to/file#fragment",
			expected: "HTTP:EXAMPLE.COM/PATH/TO/FILE",
			ok:       true,
		},
		{
			name:     "url with path and query",
			input:    "http://example.com/path/to/file?query=param#fragment",
			expected: "HTTP:EXAMPLE.COM/PATH/TO/FILE",
			ok:       true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, ok := UrlToFilename(test.input)
			if ok != test.ok {
				t.Errorf("Expected ok to be %v, but got %v", test.ok, ok)
			}
			if got != test.expected {
				t.Errorf("Expected %q, but got %q", test.expected, got)
			}
		})
	}
}
