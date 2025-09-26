package images

import (
	"fmt"
	"testing"
)

func TestAbsf32_1(t *testing.T) {
	t.Run("Testing positive float32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := absf32(1.23)
		want := float32(1.23)
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Testing negative float32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := absf32(-1.23)
		want := float32(1.23)
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})

	t.Run("Testing zero float32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := absf32(0)
		want := float32(0)
		if got != want {
			t.Errorf("got %f want %f", got, want)
		}
	})
}
