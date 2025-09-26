package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCollectIdentities_1(t *testing.T) {
	t.Run("should return error if elements are not comparable", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := collectIdentities([]any{1, 2, 3, 1})
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("should return error if arguments are not slices or arrays", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := collectIdentities(1, 2, 3)
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("should return map of unique elements", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := map[any]bool{1: true, 2: true, 3: true}
		result, err := collectIdentities([]any{1, 2, 3})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
