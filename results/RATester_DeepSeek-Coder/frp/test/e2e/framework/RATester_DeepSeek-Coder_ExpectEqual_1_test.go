package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectEqual_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectEqual", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectEqual(t, 1, 1, "1 should equal 1")
		ExpectEqual(t, 2, 1, "2 should not equal 1")
	})
}
