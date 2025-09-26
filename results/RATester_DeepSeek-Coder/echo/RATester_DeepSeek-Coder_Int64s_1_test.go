package echo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInt64s_1(t *testing.T) {
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

	t.Run("TestInt64s", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []int64
		b.Int64s("test", &dest)
		expected := []int64{1, 2, 3}
		if !reflect.DeepEqual(dest, expected) {
			t.Errorf("Expected %v, got %v", expected, dest)
		}
	})

	t.Run("TestInt64s_Empty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []int64
		b.Int64s("", &dest)
		if len(dest) != 0 {
			t.Errorf("Expected empty slice, got %v", dest)
		}
	})

	t.Run("TestInt64s_Invalid", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		var dest []int64
		b.Int64s("invalid", &dest)
		if len(dest) != 0 {
			t.Errorf("Expected empty slice, got %v", dest)
		}
	})
}
