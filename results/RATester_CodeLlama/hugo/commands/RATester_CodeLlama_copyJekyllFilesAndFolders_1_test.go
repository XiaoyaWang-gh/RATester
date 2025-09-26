package commands

import (
	"fmt"
	"testing"
)

func TestCopyJekyllFilesAndFolders_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// TODO: Setup the test case
	// TODO: Test the method
	// TODO: Verify the results
}
