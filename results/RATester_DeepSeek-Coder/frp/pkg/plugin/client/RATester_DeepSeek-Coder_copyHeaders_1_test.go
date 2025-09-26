package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCopyHeaders_1(t *testing.T) {
	t.Run("copyHeaders", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		src := make(http.Header)
		dst := make(http.Header)

		src.Set("Content-Type", "application/json")
		src.Set("Authorization", "Bearer token")

		copyHeaders(dst, src)

		if dst.Get("Content-Type") != "application/json" {
			t.Errorf("Expected 'application/json', got '%s'", dst.Get("Content-Type"))
		}

		if dst.Get("Authorization") != "Bearer token" {
			t.Errorf("Expected 'Bearer token', got '%s'", dst.Get("Authorization"))
		}
	})
}
