package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectContainElements_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectContainElements", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		actual := []int{1, 2, 3, 4, 5}
		extra := []int{1, 2}
		ExpectContainElements(t, actual, extra, "actual should contain extra elements")
	})
}
