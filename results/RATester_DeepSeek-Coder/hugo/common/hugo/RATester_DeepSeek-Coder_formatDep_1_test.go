package hugo

import (
	"fmt"
	"testing"
)

func TestFormatDep_1(t *testing.T) {
	type test struct {
		path     string
		version  string
		expected string
	}

	tests := []test{
		{
			path:     "github.com/golang/dep",
			version:  "1.2.3",
			expected: "github.com/golang/dep=\"1.2.3\"",
		},
		{
			path:     "github.com/golang/dep",
			version:  "master",
			expected: "github.com/golang/dep=\"master\"",
		},
		{
			path:     "github.com/golang/dep",
			version:  "",
			expected: "github.com/golang/dep=\"\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := formatDep(tt.path, tt.version)
			if got != tt.expected {
				t.Errorf("formatDep(%q, %q) = %q; want %q", tt.path, tt.version, got, tt.expected)
			}
		})
	}
}
