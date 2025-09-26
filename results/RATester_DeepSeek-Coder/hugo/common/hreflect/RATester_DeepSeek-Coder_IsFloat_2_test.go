package hreflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsFloat_2(t *testing.T) {
	t.Run("Testing with float32", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := IsFloat(reflect.Float32)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Testing with float64", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := IsFloat(reflect.Float64)
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Testing with int", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := IsFloat(reflect.Int)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Testing with string", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		got := IsFloat(reflect.String)
		want := false

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
