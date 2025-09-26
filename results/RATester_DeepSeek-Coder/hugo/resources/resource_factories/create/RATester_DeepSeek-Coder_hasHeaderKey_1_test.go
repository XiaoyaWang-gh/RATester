package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHasHeaderKey_1(t *testing.T) {
	t.Run("should return true when the key exists in the header", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		header := make(http.Header)
		header.Set("Content-Type", "application/json")
		got := hasHeaderKey(header, "Content-Type")
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return false when the key does not exist in the header", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		header := make(http.Header)
		header.Set("Content-Type", "application/json")
		got := hasHeaderKey(header, "Content-Length")
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
