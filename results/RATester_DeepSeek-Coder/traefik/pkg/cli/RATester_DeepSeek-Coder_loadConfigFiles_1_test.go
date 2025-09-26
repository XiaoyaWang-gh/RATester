package cli

import (
	"fmt"
	"testing"
)

func TestLoadConfigFiles_1(t *testing.T) {
	type testStruct struct {
		Test string
	}

	t.Run("Test with existing config file", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		element := &testStruct{}
		filePath, err := loadConfigFiles("existing_config_file.toml", element)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if filePath == "" {
			t.Errorf("Expected file path, got empty string")
		}
		if element.Test == "" {
			t.Errorf("Expected test field to be populated, got empty string")
		}
	})

	t.Run("Test with non-existing config file", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		element := &testStruct{}
		filePath, err := loadConfigFiles("non_existing_config_file.toml", element)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if filePath != "" {
			t.Errorf("Expected empty file path, got %s", filePath)
		}
		if element.Test != "" {
			t.Errorf("Expected test field to be empty, got %s", element.Test)
		}
	})
}
