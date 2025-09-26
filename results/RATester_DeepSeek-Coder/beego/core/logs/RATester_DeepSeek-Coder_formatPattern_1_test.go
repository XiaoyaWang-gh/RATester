package logs

import (
	"fmt"
	"testing"
)

func TestFormatPattern_1(t *testing.T) {
	t.Run("Test with string and no variable", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := formatPattern("test")
		want := "test"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test with string and variables", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := formatPattern("test %v %v", 1, 2)
		want := "test 1 2"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test with non-string and no variables", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := formatPattern(1)
		want := "1"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test with non-string and variables", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := formatPattern(1, 2, 3)
		want := "1 2 3"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Test with string that contains %", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := formatPattern("%test %v %v", 1, 2)
		want := "%test 1 2"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
