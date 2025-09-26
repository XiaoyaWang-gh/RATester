package customerrors

import (
	"fmt"
	"testing"
)

func TestIsFilteredCode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &codeCatcher{}
	cc.caughtFilteredCode = true
	if !cc.isFilteredCode() {
		t.Errorf("Expected true, got false")
	}
}
