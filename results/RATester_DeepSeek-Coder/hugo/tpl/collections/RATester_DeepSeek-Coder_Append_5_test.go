package collections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend_5(t *testing.T) {
	ns := &Namespace{}

	t.Run("should return error when less than 2 arguments are provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.Append()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should append elements from the first arguments to the last argument", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := []any{1, 2, 3, 4, 5}
		result, err := ns.Append([]any{1, 2, 3}, 4, 5)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
