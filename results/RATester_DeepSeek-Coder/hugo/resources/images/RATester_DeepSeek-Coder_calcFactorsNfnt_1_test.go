package images

import (
	"fmt"
	"testing"
)

func TestCalcFactorsNfnt_1(t *testing.T) {
	t.Run("Test with width and height both 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		scaleX, scaleY := calcFactorsNfnt(0, 0, 100.0, 100.0)
		if scaleX != 1.0 || scaleY != 1.0 {
			t.Errorf("Expected (1.0, 1.0), got (%f, %f)", scaleX, scaleY)
		}
	})

	t.Run("Test with width 0 and height not 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		scaleX, scaleY := calcFactorsNfnt(0, 50, 100.0, 100.0)
		if scaleX != 1.0 || scaleY != 2.0 {
			t.Errorf("Expected (1.0, 2.0), got (%f, %f)", scaleX, scaleY)
		}
	})

	t.Run("Test with width not 0 and height 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		scaleX, scaleY := calcFactorsNfnt(50, 0, 100.0, 100.0)
		if scaleX != 2.0 || scaleY != 1.0 {
			t.Errorf("Expected (2.0, 1.0), got (%f, %f)", scaleX, scaleY)
		}
	})

	t.Run("Test with width and height not 0", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		scaleX, scaleY := calcFactorsNfnt(50, 50, 100.0, 100.0)
		if scaleX != 2.0 || scaleY != 2.0 {
			t.Errorf("Expected (2.0, 2.0), got (%f, %f)", scaleX, scaleY)
		}
	})
}
