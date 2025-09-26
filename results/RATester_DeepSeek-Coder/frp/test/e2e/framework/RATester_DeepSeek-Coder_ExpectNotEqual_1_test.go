package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectNotEqual_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectNotEqual", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectNotEqual(t, 1, 2, "1 should not equal 2")
		ExpectNotEqual(t, "hello", "world", "hello should not equal world")
		ExpectNotEqual(t, []int{1, 2, 3}, []int{1, 2, 3}, "slices should not be equal")
	})
}
