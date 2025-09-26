package config

import (
	"fmt"
	"testing"
)

func TestDetectLegacyINIFormat_1(t *testing.T) {
	t.Run("Test with valid INI content", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		content := []byte(`
		[common]
		key1 = value1
		key2 = value2
		`)
		result := DetectLegacyINIFormat(content)
		if !result {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("Test with invalid INI content", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		content := []byte(`
		key1 = value1
		key2 = value2
		`)
		result := DetectLegacyINIFormat(content)
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	t.Run("Test with empty content", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		content := []byte("")
		result := DetectLegacyINIFormat(content)
		if result {
			t.Errorf("Expected false, got true")
		}
	})
}
