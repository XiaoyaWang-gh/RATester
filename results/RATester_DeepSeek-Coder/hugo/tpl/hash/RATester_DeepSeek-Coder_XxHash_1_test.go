package hash

import (
	"fmt"
	"testing"
)

func TestXxHash_1(t *testing.T) {
	t.Parallel()

	ns := &Namespace{}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		result, err := ns.XxHash("test")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		expected := "9e107d9d372bb682"
		if result != expected {
			t.Errorf("Expected %q, got %q", expected, result)
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.XxHash(make(chan int))
		if err == nil {
			t.Errorf("Expected an error, got nil")
		}
	})
}
