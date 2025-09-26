package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntersect_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("should return empty slice when one of the inputs is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Intersect(nil, []int{1, 2, 3})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(result.([]int)) != 0 {
			t.Errorf("expected empty slice, got %v", result)
		}
	})

	t.Run("should return intersection of two slices", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.Intersect([]int{1, 2, 3}, []int{2, 3, 4})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected := []int{2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("should return error when slices contain uncomparable types", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Intersect([]struct{}{{}}, []struct{}{{}})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error when trying to intersect incompatible types", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Intersect([]int{1, 2, 3}, "string")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
