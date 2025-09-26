package task

import (
	"fmt"
	"testing"
)

func TestParseIntOrName_1(t *testing.T) {
	names := map[string]uint{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	t.Run("Test with named integer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := parseIntOrName("one", names)
		want := uint(1)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Test with integer string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := parseIntOrName("2", names)
		want := uint(2)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Test with unnamed integer", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := parseIntOrName("4", nil)
		want := uint(4)
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
