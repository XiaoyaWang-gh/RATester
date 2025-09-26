package paths

import (
	"fmt"
	"testing"
)

func TestFilename_2(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		expected string
	}{
		{"Test case 1", "test.txt", "test"},
		{"Test case 2", "test.pdf", "test"},
		{"Test case 3", "test", "test"},
		{"Test case 4", "test.txt.gz", "test.txt"},
		{"Test case 5", "test.txt.tar.gz", "test.txt.tar"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := Filename(test.in)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
