package echo

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestFilepathOrContent_1(t *testing.T) {
	t.Run("Test with valid file path", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		filePath := "testdata/test.txt"
		content, err := filepathOrContent(filePath)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expectedContent, _ := os.ReadFile(filePath)
		if !bytes.Equal(content, expectedContent) {
			t.Errorf("Expected content %v, got %v", expectedContent, content)
		}
	})

	t.Run("Test with valid byte slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		content := []byte("test content")
		result, err := filepathOrContent(content)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !bytes.Equal(result, content) {
			t.Errorf("Expected content %v, got %v", content, result)
		}
	})

	t.Run("Test with invalid input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := filepathOrContent(123)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
