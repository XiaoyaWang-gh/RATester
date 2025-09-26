package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectNotContainElements_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectNotContainElements", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		actual := []int{1, 2, 3, 4, 5}
		extra := []int{2, 4}
		ExpectNotContainElements(actual, extra, "actual should not contain extra elements")
	})
}
