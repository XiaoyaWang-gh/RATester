package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInt16s_1(t *testing.T) {
	b := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "1"
			}
			return ""
		},
		ValuesFunc: func(sourceParam string) []string {
			if sourceParam == "test" {
				return []string{"1", "2", "3"}
			}
			return nil
		},
	}

	t.Run("TestInt16s", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []int16
		b.Int16s("test", &dest)
		expected := []int16{1, 2, 3}
		if !reflect.DeepEqual(dest, expected) {
			t.Errorf("Expected %v, got %v", expected, dest)
		}
	})

	t.Run("TestInt16s_Empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []int16
		b.Int16s("not_exist", &dest)
		if len(dest) != 0 {
			t.Errorf("Expected empty slice, got %v", dest)
		}
	})
}
