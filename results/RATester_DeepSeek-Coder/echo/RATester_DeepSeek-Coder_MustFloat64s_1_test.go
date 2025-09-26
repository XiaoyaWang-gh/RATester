package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustFloat64s_1(t *testing.T) {
	t.Run("Test MustFloat64s with valid values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValuesFunc: func(sourceParam string) []string {
				return []string{"1.1", "2.2", "3.3"}
			},
		}
		var dest []float64
		b.MustFloat64s("param", &dest)
		expected := []float64{1.1, 2.2, 3.3}
		if !reflect.DeepEqual(dest, expected) {
			t.Errorf("Expected %v, got %v", expected, dest)
		}
	})

	t.Run("Test MustFloat64s with invalid values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValuesFunc: func(sourceParam string) []string {
				return []string{"invalid", "2.2", "3.3"}
			},
		}
		var dest []float64
		err := b.MustFloat64s("param", &dest)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("Test MustFloat64s with no values", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		b := &ValueBinder{
			ValuesFunc: func(sourceParam string) []string {
				return []string{}
			},
		}
		var dest []float64
		err := b.MustFloat64s("param", &dest)
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})
}
