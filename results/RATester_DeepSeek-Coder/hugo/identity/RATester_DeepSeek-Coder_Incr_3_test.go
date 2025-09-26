package identity

import (
	"fmt"
	"testing"
)

func TestIncr_3(t *testing.T) {
	incr := &IncrementByOne{}

	t.Run("Incr", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := incr.Incr()
		want := 1

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Incr again", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := incr.Incr()
		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
