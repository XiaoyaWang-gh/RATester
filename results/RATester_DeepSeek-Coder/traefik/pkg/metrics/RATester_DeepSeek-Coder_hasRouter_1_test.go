package metrics

import (
	"fmt"
	"testing"
)

func TestHasRouter_1(t *testing.T) {
	d := &dynamicConfig{
		routers: map[string]bool{
			"router1": true,
			"router2": true,
		},
	}

	t.Run("hasRouter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("returns true when router exists", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := d.hasRouter("router1")
			want := true

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})

		t.Run("returns false when router does not exist", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := d.hasRouter("router3")
			want := false

			if got != want {
				t.Errorf("got %t, want %t", got, want)
			}
		})
	})
}
