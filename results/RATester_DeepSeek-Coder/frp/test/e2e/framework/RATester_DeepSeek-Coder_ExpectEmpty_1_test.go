package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectEmpty_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectEmpty", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Run("should pass when actual is empty", func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ExpectEmpty([]int{})
		})

		t.Run("should fail when actual is not empty", func(t *testing.T) {

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
			ExpectEmpty([]int{1})
		})
	})
}
