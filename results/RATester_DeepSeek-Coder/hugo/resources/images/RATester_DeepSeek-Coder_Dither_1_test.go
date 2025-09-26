package images

import (
	"fmt"
	"testing"
)

func TestDither_1(t *testing.T) {
	t.Run("Test with valid options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		f := &Filters{}
		filter := f.Dither([]any{
			map[string]any{
				"Colors":     []any{"000000ff", "ffffffff"},
				"Method":     "floydsteinberg",
				"Serpentine": true,
				"Strength":   1.0,
			},
		})
		if filter == nil {
			t.Errorf("Expected filter to be not nil")
		}
	})

	t.Run("Test with invalid options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		f := &Filters{}
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic with invalid options")
			}
		}()
		f.Dither([]any{
			map[string]any{
				"Colors":     []any{"000000ff", "invalid"},
				"Method":     "floydsteinberg",
				"Serpentine": true,
				"Strength":   1.0,
			},
		})
	})

	t.Run("Test with no options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		f := &Filters{}
		filter := f.Dither()
		if filter == nil {
			t.Errorf("Expected filter to be not nil")
		}
	})
}
