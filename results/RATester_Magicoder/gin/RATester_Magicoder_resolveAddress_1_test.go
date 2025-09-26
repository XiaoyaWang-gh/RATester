package gin

import (
	"fmt"
	"os"
	"testing"
)

func TestresolveAddress_1(t *testing.T) {
	t.Run("EmptyAddress", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("PORT", "8080")
		defer os.Unsetenv("PORT")
		addr := resolveAddress([]string{})
		if addr != ":8080" {
			t.Errorf("Expected :8080, got %s", addr)
		}
	})

	t.Run("SingleAddress", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		addr := resolveAddress([]string{":8080"})
		if addr != ":8080" {
			t.Errorf("Expected :8080, got %s", addr)
		}
	})

	t.Run("TooManyAddresses", func(t *testing.T) {

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
		resolveAddress([]string{":8080", ":8081"})
	})
}
