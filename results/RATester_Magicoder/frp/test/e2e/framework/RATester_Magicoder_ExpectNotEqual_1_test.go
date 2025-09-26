package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectNotEqual_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gomega.RegisterTestingT(t)

	// Test case 1: Actual and Extra are equal
	actual := 1
	extra := 1
	ExpectNotEqual(t, actual, extra, "Actual and Extra should not be equal")

	// Test case 2: Actual and Extra are not equal
	actual = 2
	extra = 1
	ExpectNotEqual(t, actual, extra, "Actual and Extra should be equal")
}
