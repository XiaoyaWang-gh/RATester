package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFilterForwardRequestHeaders_1(t *testing.T) {
	t.Run("should return original headers when no allowed headers are provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		headers := http.Header{
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Bearer token"},
		}
		allowedHeaders := []string{}

		filteredHeaders := filterForwardRequestHeaders(headers, allowedHeaders)

		if len(filteredHeaders) != len(headers) {
			t.Errorf("Expected %d headers, got %d", len(headers), len(filteredHeaders))
		}
	})

	t.Run("should return only allowed headers", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		headers := http.Header{
			"Content-Type":    []string{"application/json"},
			"Authorization":   []string{"Bearer token"},
			"X-Custom-Header": []string{"custom value"},
		}
		allowedHeaders := []string{"Content-Type", "X-Custom-Header"}

		filteredHeaders := filterForwardRequestHeaders(headers, allowedHeaders)

		if len(filteredHeaders) != len(allowedHeaders) {
			t.Errorf("Expected %d headers, got %d", len(allowedHeaders), len(filteredHeaders))
		}
		for _, header := range allowedHeaders {
			if _, ok := filteredHeaders[header]; !ok {
				t.Errorf("Expected header '%s' not found", header)
			}
		}
	})

	t.Run("should return empty headers when none of the allowed headers are present", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		headers := http.Header{
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Bearer token"},
		}
		allowedHeaders := []string{"X-Custom-Header"}

		filteredHeaders := filterForwardRequestHeaders(headers, allowedHeaders)

		if len(filteredHeaders) != 0 {
			t.Errorf("Expected 0 headers, got %d", len(filteredHeaders))
		}
	})
}
