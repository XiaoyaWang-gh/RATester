package commands

import (
	"fmt"
	"testing"
)

func TestgetDirList_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hugoBuilder := &hugoBuilder{
		// Initialize the fields of hugoBuilder
	}

	dirList, err := hugoBuilder.getDirList()

	if err != nil {
		t.Errorf("getDirList() returned an error: %v", err)
	}

	if len(dirList) == 0 {
		t.Error("getDirList() returned an empty list")
	}

	// Add more test cases as needed
}
