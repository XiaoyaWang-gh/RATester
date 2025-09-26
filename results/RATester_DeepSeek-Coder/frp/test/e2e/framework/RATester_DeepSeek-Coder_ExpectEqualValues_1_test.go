package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectEqualValues_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("TestExpectEqualValues", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectEqualValues(t, 1, 1, "1 should equal 1")
		ExpectEqualValues(t, "hello", "hello", "hello should equal hello")
		ExpectEqualValues(t, []int{1, 2, 3}, []int{1, 2, 3}, "[]int{1, 2, 3} should equal []int{1, 2, 3}")
	})
}
