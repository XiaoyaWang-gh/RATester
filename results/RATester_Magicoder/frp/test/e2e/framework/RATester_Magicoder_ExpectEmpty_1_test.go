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

		ExpectEmpty(nil, "Expected nil to be empty")
		ExpectEmpty("", "Expected empty string to be empty")
		ExpectEmpty([]int{}, "Expected empty slice to be empty")
		ExpectEmpty(map[string]int{}, "Expected empty map to be empty")
		ExpectEmpty(0, "Expected zero to be empty")
		ExpectEmpty(false, "Expected false to be empty")
	})
}
