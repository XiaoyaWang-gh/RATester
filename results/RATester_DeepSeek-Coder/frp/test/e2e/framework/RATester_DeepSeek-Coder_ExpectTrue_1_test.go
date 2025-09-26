package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectTrue_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("ExpectTrue should pass when actual is true", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectTrue(true)
	})

	t.Run("ExpectTrue should fail when actual is false", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("ExpectTrue should panic when actual is false")
			}
		}()
		ExpectTrue(false)
	})
}
