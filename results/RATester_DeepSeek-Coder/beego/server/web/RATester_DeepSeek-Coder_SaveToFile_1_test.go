package web

import (
	"fmt"
	"os"
	"testing"
)

func TestSaveToFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Data: make(map[interface{}]interface{}),
	}

	fromFile := "testdata/test.txt"
	toFile := "testdata/test_copy.txt"

	err := ctrl.SaveToFile(fromFile, toFile)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Check if the file was copied correctly
	_, err = os.Stat(toFile)
	if os.IsNotExist(err) {
		t.Errorf("Expected file %s to exist, but it does not", toFile)
	}

	// Clean up
	err = os.Remove(toFile)
	if err != nil {
		t.Errorf("Failed to remove file %s: %v", toFile, err)
	}
}
