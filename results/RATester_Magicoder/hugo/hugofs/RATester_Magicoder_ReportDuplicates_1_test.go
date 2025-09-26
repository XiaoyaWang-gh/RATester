package hugofs

import (
	"fmt"
	"testing"
)

func TestReportDuplicates_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &createCountingFs{
		fileCount: map[string]int{
			"file1": 2,
			"file2": 1,
			"file3": 3,
		},
	}

	expected := "file1 (2), file3 (3)"
	result := fs.ReportDuplicates()

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
