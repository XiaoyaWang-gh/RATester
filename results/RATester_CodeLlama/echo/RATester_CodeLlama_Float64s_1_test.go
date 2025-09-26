package echo

import (
	"fmt"
	"testing"
)

func TestFloat64s_1(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := &[]float64{}

		b := &ValueBinder{}
		b.Float64s(sourceParam, dest)

		if len(*dest) != 0 {
			t.Errorf("expected %d, got %d", 0, len(*dest))
		}
	})

	t.Run("failure", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		sourceParam := "sourceParam"
		dest := &[]float64{}

		b := &ValueBinder{}
		b.Float64s(sourceParam, dest)

		if len(*dest) != 0 {
			t.Errorf("expected %d, got %d", 0, len(*dest))
		}
	})
}
