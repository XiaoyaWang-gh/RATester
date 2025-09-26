package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFloat64s_1(t *testing.T) {
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "1.23"
			}
			return ""
		},
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "test" {
				return []string{"1.23", "4.56"}
			}
			return nil
		},
	}

	t.Run("TestFloat64s", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []float64
		b.Float64s("test", &dest)
		expected := []float64{1.23, 4.56}
		if !reflect.DeepEqual(dest, expected) {
			t.Errorf("Expected %v, got %v", expected, dest)
		}
	})
}
