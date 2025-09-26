package framework

import (
	"fmt"
	"testing"

	"github.com/bsm/gomega"
)

func TestExpectEqual_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	gomega.RegisterTestingT(t)

	// Test case 1: When actual and extra are equal
	ExpectEqual(t, 1, 1, "Test case 1 failed")

	// Test case 2: When actual and extra are not equal
	ExpectEqual(t, 1, 2, "Test case 2 failed")
}
