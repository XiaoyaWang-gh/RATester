package data

import (
	"fmt"
	"testing"
)

func TestToURLAndHeaders_1(t *testing.T) {
	t.Run("Test with empty urlParts", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		url, headers := toURLAndHeaders([]any{})
		if url != "" || headers != nil {
			t.Errorf("Expected url to be empty and headers to be nil, got %v and %v", url, headers)
		}
	})

	t.Run("Test with urlParts and headers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		urlParts := []any{"http://example.com", map[string]any{"Content-Type": "application/json"}}
		url, headers := toURLAndHeaders(urlParts)
		if url != "http://example.com" || headers["Content-Type"] != "application/json" {
			t.Errorf("Expected url to be 'http://example.com' and headers to be {'Content-Type': 'application/json'}, got %v and %v", url, headers)
		}
	})

	t.Run("Test with urlParts and no headers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		urlParts := []any{"http://example.com"}
		url, headers := toURLAndHeaders(urlParts)
		if url != "http://example.com" || headers != nil {
			t.Errorf("Expected url to be 'http://example.com' and headers to be nil, got %v and %v", url, headers)
		}
	})
}
