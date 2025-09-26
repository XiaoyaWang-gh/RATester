package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadSyncMapFromFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filePath := "test.txt"
	expectedContent := "test1\ntest2\ntest3"
	err := ioutil.WriteFile(filePath, []byte(expectedContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filePath)

	var actualContent []string
	loadSyncMapFromFile(filePath, func(value string) {
		actualContent = append(actualContent, value)
	})

	if len(actualContent) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(actualContent))
	}

	expectedLines := []string{"test1", "test2", "test3"}
	for i, line := range actualContent {
		if line != expectedLines[i] {
			t.Errorf("Expected line %d to be '%s', got '%s'", i+1, expectedLines[i], line)
		}
	}
}
