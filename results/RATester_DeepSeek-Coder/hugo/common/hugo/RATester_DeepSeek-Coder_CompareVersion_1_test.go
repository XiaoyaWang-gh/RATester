package hugo

import (
	"fmt"
	"testing"
)

func TestCompareVersion_1(t *testing.T) {
	t.Run("Test with valid version", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		version := "1.2.3"
		result := CompareVersion(version)
		if result != 0 {
			t.Errorf("Expected 0, got %d", result)
		}
	})

	t.Run("Test with invalid version", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		version := "1.2.3.4"
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		CompareVersion(version)
	})

	t.Run("Test with non-string input", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		version := 123
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		CompareVersion(version)
	})
}
