package constraints

import (
	"fmt"
	"testing"
)

func TestTagFn_1(t *testing.T) {
	name := "test"
	tagFn := tagFn(name)

	t.Run("should return true when tag exists", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tags := []string{name, "other"}
		if !tagFn(tags) {
			t.Errorf("Expected true, got false")
		}
	})

	t.Run("should return false when tag does not exist", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		tags := []string{"other"}
		if tagFn(tags) {
			t.Errorf("Expected false, got true")
		}
	})
}
