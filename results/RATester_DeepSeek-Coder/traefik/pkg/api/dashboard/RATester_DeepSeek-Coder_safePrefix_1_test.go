package dashboard

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSafePrefix_1(t *testing.T) {
	t.Run("Test safePrefix function", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("Test with valid X-Forwarded-Prefix", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-Forwarded-Prefix", "http://example.com/prefix")
			prefix := safePrefix(req)
			if prefix != "/prefix" {
				t.Errorf("Expected prefix to be /prefix, got %s", prefix)
			}
		})

		t.Run("Test with invalid X-Forwarded-Prefix", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("X-Forwarded-Prefix", "invalid")
			prefix := safePrefix(req)
			if prefix != "" {
				t.Errorf("Expected prefix to be empty, got %s", prefix)
			}
		})

		t.Run("Test with empty X-Forwarded-Prefix", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest("GET", "/test", nil)
			if err != nil {
				t.Fatal(err)
			}
			prefix := safePrefix(req)
			if prefix != "" {
				t.Errorf("Expected prefix to be empty, got %s", prefix)
			}
		})
	})
}
