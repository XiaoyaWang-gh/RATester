package orm

import (
	"fmt"
	"testing"
)

func TestOrNot_1(t *testing.T) {
	t.Run("test OrNot with empty expr", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		c := Condition{}
		c.OrNot("", "test")
	})

	t.Run("test OrNot with empty args", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		c := Condition{}
		c.OrNot("test", "")
	})

	t.Run("test OrNot with valid args", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c := Condition{}
		result := c.OrNot("test", "test")
		if result == nil {
			t.Errorf("Expected a non-nil result")
		}
		if len(result.params) != 1 {
			t.Errorf("Expected 1 condition, got %d", len(result.params))
		}
		if result.params[0].exprs[0] != "test" {
			t.Errorf("Expected 'test', got %s", result.params[0].exprs[0])
		}
		if len(result.params[0].args) != 1 {
			t.Errorf("Expected 1 arg, got %d", len(result.params[0].args))
		}
		if result.params[0].args[0] != "test" {
			t.Errorf("Expected 'test', got %s", result.params[0].args[0])
		}
		if !result.params[0].isOr {
			t.Errorf("Expected isOr to be true")
		}
		if !result.params[0].isNot {
			t.Errorf("Expected isNot to be true")
		}
	})
}
