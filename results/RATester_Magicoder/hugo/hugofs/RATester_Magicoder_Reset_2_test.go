package hugofs

import (
	"fmt"
	"testing"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &createCountingFs{
		fileCount: map[string]int{
			"file1": 1,
			"file2": 2,
		},
	}

	fs.Reset()

	if len(fs.fileCount) != 0 {
		t.Errorf("Expected fileCount to be empty after Reset, but got %v", fs.fileCount)
	}
}
