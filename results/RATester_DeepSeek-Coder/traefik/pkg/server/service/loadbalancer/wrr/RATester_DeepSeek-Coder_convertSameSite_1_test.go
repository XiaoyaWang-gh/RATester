package wrr

import (
	"fmt"
	"net/http"
	"testing"
)

func TestConvertSameSite_1(t *testing.T) {
	t.Run("TestNone", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := convertSameSite("none")
		if result != http.SameSiteNoneMode {
			t.Errorf("Expected http.SameSiteNoneMode, got %v", result)
		}
	})

	t.Run("TestLax", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := convertSameSite("lax")
		if result != http.SameSiteLaxMode {
			t.Errorf("Expected http.SameSiteLaxMode, got %v", result)
		}
	})

	t.Run("TestStrict", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := convertSameSite("strict")
		if result != http.SameSiteStrictMode {
			t.Errorf("Expected http.SameSiteStrictMode, got %v", result)
		}
	})

	t.Run("TestDefault", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result := convertSameSite("default")
		if result != http.SameSiteDefaultMode {
			t.Errorf("Expected http.SameSiteDefaultMode, got %v", result)
		}
	})
}
