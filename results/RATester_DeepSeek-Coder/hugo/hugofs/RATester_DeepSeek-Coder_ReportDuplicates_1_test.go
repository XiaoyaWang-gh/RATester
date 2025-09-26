package hugofs

import (
	"fmt"
	"testing"
)

func TestReportDuplicates_1(t *testing.T) {
	tests := []struct {
		name     string
		files    []string
		expected string
	}{
		{
			name:     "no duplicates",
			files:    []string{"file1", "file2", "file3"},
			expected: "",
		},
		{
			name:     "one duplicate",
			files:    []string{"file1", "file1", "file2"},
			expected: "file1 (2)",
		},
		{
			name:     "multiple duplicates",
			files:    []string{"file1", "file1", "file2", "file2", "file3", "file3"},
			expected: "file1 (2), file2 (2), file3 (2)",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			fs := &createCountingFs{
				fileCount: make(map[string]int),
			}

			for _, file := range test.files {
				fs.onCreate(file)
			}

			result := fs.ReportDuplicates()
			if result != test.expected {
				t.Errorf("Expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
