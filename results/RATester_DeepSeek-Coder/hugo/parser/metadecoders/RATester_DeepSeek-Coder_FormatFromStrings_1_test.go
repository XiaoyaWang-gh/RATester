package metadecoders

import (
	"fmt"
	"testing"
)

func TestFormatFromStrings_1(t *testing.T) {
	t.Run("should return the first valid format", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := FormatFromStrings("invalid", "valid", "alsoInvalid")
		want := Format("valid")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("should return empty string if no valid format", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := FormatFromStrings("invalid", "alsoInvalid")
		want := Format("")

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
