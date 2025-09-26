package glob

import (
	"fmt"
	"testing"
)

func TestGetGlob_2(t *testing.T) {
	t.Run("should return a valid glob", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		g, err := GetGlob("*")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if g == nil {
			t.Errorf("expected a valid glob, got nil")
		}
	})

	t.Run("should return an error for an invalid pattern", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := GetGlob("[")
		if err == nil {
			t.Errorf("expected an error, got nil")
		}
	})
}
