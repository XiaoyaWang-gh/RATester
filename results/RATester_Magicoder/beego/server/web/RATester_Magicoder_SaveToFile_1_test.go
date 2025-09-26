package web

import (
	"fmt"
	"testing"
)

func TestSaveToFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}
	ctrl.Init(nil, "TestController", "TestAction", nil)

	fromFile := "testdata/fromFile.txt"
	toFile := "testdata/toFile.txt"

	err := ctrl.SaveToFile(fromFile, toFile)
	if err != nil {
		t.Errorf("SaveToFile failed: %v", err)
	}

	// Add assertions to check if the file was correctly saved
	// For example, you can read the content of the saved file and compare it with the content of the original file
}
