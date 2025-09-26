package framework

import (
	"errors"
	"fmt"
	"testing"
)

func TestExpectNoError_1(t *testing.T) {
	t.Run("NoError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectNoError(nil)
	})

	t.Run("Error", func(t *testing.T) {

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
		ExpectNoError(errors.New("error"))
	})
}
