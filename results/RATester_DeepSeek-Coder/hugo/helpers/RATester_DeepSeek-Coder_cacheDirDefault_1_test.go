package helpers

import (
	"fmt"
	"os"
	"testing"
)

func TestCacheDirDefault_1(t *testing.T) {

	t.Run("Test with non-empty cacheDir", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cacheDir := "/tmp"
		expected := "/tmp/"
		result := cacheDirDefault(cacheDir)
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Test with empty cacheDir", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		cacheDir := ""
		expected := ""
		result := cacheDirDefault(cacheDir)
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Test with NETLIFY=true, PULL_REQUEST=, DEPLOY_PRIME_URL=", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		os.Setenv("NETLIFY", "true")
		os.Setenv("PULL_REQUEST", "")
		os.Setenv("DEPLOY_PRIME_URL", "")
		cacheDir := ""
		expected := "/opt/build/cache/hugo_cache/"
		result := cacheDirDefault(cacheDir)
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
		os.Unsetenv("NETLIFY")
		os.Unsetenv("PULL_REQUEST")
		os.Unsetenv("DEPLOY_PRIME_URL")
	})
}
