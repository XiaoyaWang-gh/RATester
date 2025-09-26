package page

import (
	"fmt"
	"testing"
)

func TestPathFile_1(t *testing.T) {
	tests := []struct {
		name     string
		builder  *pagePathBuilder
		expected string
	}{
		{
			name: "Test case 1",
			builder: &pagePathBuilder{
				els:        []string{"test", "case", "1"},
				prefixPath: "",
			},
			expected: "/test/case/1",
		},
		{
			name: "Test case 2",
			builder: &pagePathBuilder{
				els:        []string{"test", "case", "2"},
				prefixPath: "prefix/",
			},
			expected: "/prefix/test/case/2",
		},
		{
			name: "Test case 3",
			builder: &pagePathBuilder{
				els:        []string{"test", "case", "3"},
				prefixPath: "/prefix/",
			},
			expected: "/prefix/test/case/3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.builder.PathFile()
			if result != test.expected {
				t.Errorf("Expected %s, but got %s", test.expected, result)
			}
		})
	}
}
