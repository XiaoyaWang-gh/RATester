package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse_1(t *testing.T) {
	ns := &Namespace{}

	t.Run("should return error if argument is not a slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Reverse("not a slice")
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return nil if argument is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		res, err := ns.Reverse(nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if res != nil {
			t.Errorf("expected nil, got %v", res)
		}
	})

	t.Run("should reverse a slice", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		res, err := ns.Reverse([]int{1, 2, 3, 4, 5})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected := []int{5, 4, 3, 2, 1}
		if !reflect.DeepEqual(res, expected) {
			t.Errorf("expected %v, got %v", expected, res)
		}
	})
}
