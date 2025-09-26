package orm

import (
	"fmt"
	"testing"
)

func TestAnd_1(t *testing.T) {
	t.Run("TestAnd_EmptyExpr", func(t *testing.T) {

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
		c.And("", nil)
	})

	t.Run("TestAnd_EmptyArgs", func(t *testing.T) {

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
		c.And("test", nil)
	})

	t.Run("TestAnd_Success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		c := Condition{}
		res := c.And("test", "test")
		if res == nil {
			t.Errorf("Expected a non-nil result")
		}
	})
}
