package media

import (
	"fmt"
	"testing"
)

func TestIsContentSuffix_1(t *testing.T) {
	contentTypes := ContentTypes{
		extensionSet: map[string]bool{
			".html": true,
			".md":   true,
			".txt":  true,
		},
	}

	t.Run("returns true for existing suffix", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := contentTypes.IsContentSuffix(".md")
		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("returns false for non-existing suffix", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := contentTypes.IsContentSuffix(".doc")
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}
