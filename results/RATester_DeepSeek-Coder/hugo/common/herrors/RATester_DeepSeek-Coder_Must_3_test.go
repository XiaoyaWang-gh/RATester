package herrors

import (
	"errors"
	"fmt"
	"testing"
)

func TestMust_3(t *testing.T) {
	t.Run("should not panic when err is nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Expected no panic, but got: %v", r)
			}
		}()
		Must(nil)
	})

	t.Run("should panic when err is not nil", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic, but got nothing")
			}
		}()
		Must(errors.New("test error"))
	})
}
