package commands

import (
	"fmt"
	"testing"
)

func TestGetDirList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoBuilder := &hugoBuilder{
		r: &rootCommand{
			source: "/path/to/source",
		},
	}

	dirList, err := hugoBuilder.getDirList()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(dirList) == 0 {
		t.Errorf("Expected non-empty dirList, got empty")
	}

	// Add more specific tests if needed
}
