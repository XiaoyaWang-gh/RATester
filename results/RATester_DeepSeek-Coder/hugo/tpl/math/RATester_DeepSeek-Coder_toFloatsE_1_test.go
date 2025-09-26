package math

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToFloatsE_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("single value", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		v := float64(1.23)
		expected := []float64{v}
		floats, isSlice, err := ns.toFloatsE(v)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if isSlice {
			t.Errorf("Expected isSlice to be false, got true")
		}
		if !reflect.DeepEqual(floats, expected) {
			t.Errorf("Expected %v, got %v", expected, floats)
		}
	})

	t.Run("slice of values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		v := []any{1.23, 4.56, 7.89}
		expected := []float64{1.23, 4.56, 7.89}
		floats, isSlice, err := ns.toFloatsE(v)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !isSlice {
			t.Errorf("Expected isSlice to be true, got false")
		}
		if !reflect.DeepEqual(floats, expected) {
			t.Errorf("Expected %v, got %v", expected, floats)
		}
	})

	t.Run("unsupported type", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		v := "unsupported"
		_, _, err := ns.toFloatsE(v)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
