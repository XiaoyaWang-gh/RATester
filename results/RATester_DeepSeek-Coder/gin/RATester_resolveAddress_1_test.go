package gin

import (
	"fmt"
	"os"
	"testing"
)

func TestResolveAddress_1(t *testing.T) {
	t.Run("no arguments, PORT defined", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("PORT", "8080")
		defer os.Unsetenv("PORT")
		result := resolveAddress(nil)
		if result != ":8080" {
			t.Errorf("Expected :8080, got %s", result)
		}
	})

	t.Run("no arguments, PORT undefined", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Unsetenv("PORT")
		result := resolveAddress(nil)
		if result != ":8080" {
			t.Errorf("Expected :8080, got %s", result)
		}
	})

	t.Run("single argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := resolveAddress([]string{"localhost:8080"})
		if result != "localhost:8080" {
			t.Errorf("Expected localhost:8080, got %s", result)
		}
	})

	t.Run("multiple arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		resolveAddress([]string{"localhost:8080", "localhost:8081"})
	})
}
