package framework

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectErrorWithOffset_1(t *testing.T) {
	gomega.RegisterTestingT(t)

	t.Run("Test with nil error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectErrorWithOffset(0, nil, "Expected error but got nil")
	})

	t.Run("Test with non-nil error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ExpectErrorWithOffset(0, errors.New("Test error"), "Expected error but got nil")
	})
}
