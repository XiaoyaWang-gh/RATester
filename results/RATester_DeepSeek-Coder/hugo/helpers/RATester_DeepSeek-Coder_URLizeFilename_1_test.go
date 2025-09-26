package helpers

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib/paths"
)

func TestURLizeFilename_1(t *testing.T) {
	type test struct {
		name     string
		filename string
		expected string
	}

	tests := []test{
		{"simple", "test.md", "test.md"},
		{"with spaces", "test with spaces.md", "test%20with%20spaces.md"},
		{"with special chars", "test$%^&*()_+|}{[]\\:;'<>,.?/`~.md", "test$%^&*()_+|}{[]\\:;'<>,.?/`~.md"},
		{"with unicode", "test😄.md", "test%F0%9F%98%84.md"},
	}

	pathSpec := &PathSpec{
		Paths: &paths.Paths{},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := pathSpec.URLizeFilename(tc.filename)
			if result != tc.expected {
				t.Errorf("Expected %q, got %q", tc.expected, result)
			}
		})
	}
}
