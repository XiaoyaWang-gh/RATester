package htesting

import (
	"fmt"
	"os"
	"testing"
)

func TestSupportsAll_1(t *testing.T) {
	t.Run("Test when GitHub Action is detected", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("GITHUB_ACTIONS", "true")
		defer os.Unsetenv("GITHUB_ACTIONS")
		if !SupportsAll() {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("Test when CI_LOCAL is detected", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("CI_LOCAL", "true")
		defer os.Unsetenv("CI_LOCAL")
		if !SupportsAll() {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("Test when neither GitHub Action nor CI_LOCAL is detected", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if SupportsAll() {
			t.Errorf("Expected false, got true")
		}
	})
}
