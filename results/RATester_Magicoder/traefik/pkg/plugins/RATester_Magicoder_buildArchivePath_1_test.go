package plugins

import (
	"fmt"
	"testing"
)

func TestBuildArchivePath_1(t *testing.T) {
	client := &Client{
		archives: "/tmp/archives",
	}

	testCases := []struct {
		name     string
		pName    string
		pVersion string
		expected string
	}{
		{
			name:     "simple",
			pName:    "simple",
			pVersion: "1.0.0",
			expected: "/tmp/archives/simple/1.0.0.zip",
		},
		{
			name:     "with_slash",
			pName:    "with/slash",
			pVersion: "1.0.0",
			expected: "/tmp/archives/with/slash/1.0.0.zip",
		},
		{
			name:     "with_dot",
			pName:    "with.dot",
			pVersion: "1.0.0",
			expected: "/tmp/archives/with.dot/1.0.0.zip",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := client.buildArchivePath(tc.pName, tc.pVersion)
			if result != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, result)
			}
		})
	}
}
