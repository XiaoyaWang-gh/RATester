package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectConsistOf_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectConsistOf", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		actual := []int{1, 2, 3}
		extra := []int{1, 2, 3}
		ExpectConsistOf(t, actual, extra, "actual should consist of extra")
	})

	t.Run("ExpectConsistOf with different order", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		actual := []int{3, 2, 1}
		extra := []int{1, 2, 3}
		ExpectConsistOf(t, actual, extra, "actual should consist of extra")
	})

	t.Run("ExpectConsistOf with extra elements", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		actual := []int{1, 2, 3, 4, 5}
		extra := []int{1, 2, 3}
		ExpectConsistOf(t, actual, extra, "actual should consist of extra")
	})

	t.Run("ExpectConsistOf with missing elements", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		actual := []int{1, 2}
		extra := []int{1, 2, 3}
		ExpectConsistOf(t, actual, extra, "actual should consist of extra")
	})
}
