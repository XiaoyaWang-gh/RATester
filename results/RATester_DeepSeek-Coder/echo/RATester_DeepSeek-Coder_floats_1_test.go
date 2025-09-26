package echo

import (
	"fmt"
	"testing"
)

func TestFloats_1(t *testing.T) {
	b := &ValueBinder{}
	values := []string{"1.1", "2.2", "3.3"}

	t.Run("Test floats with []float64 dest", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []float64
		b.floats("sourceParam", values, &dest)
		if len(dest) != len(values) {
			t.Errorf("Expected len %d, got %d", len(values), len(dest))
		}
		for i, v := range dest {
			if v != float64(i)+1.1 {
				t.Errorf("Expected %f, got %f", float64(i)+1.1, v)
			}
		}
	})

	t.Run("Test floats with []float32 dest", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []float32
		b.floats("sourceParam", values, &dest)
		if len(dest) != len(values) {
			t.Errorf("Expected len %d, got %d", len(values), len(dest))
		}
		for i, v := range dest {
			if v != float32(i)+1.1 {
				t.Errorf("Expected %f, got %f", float32(i)+1.1, v)
			}
		}
	})
}
