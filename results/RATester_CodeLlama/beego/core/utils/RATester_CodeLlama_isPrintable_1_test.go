package utils

import (
	"fmt"
	"testing"
)

func TestIsPrintable_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var c byte
	if isPrintable(c) {
		t.Errorf("isPrintable(%c) = true, want false", c)
	}
}
