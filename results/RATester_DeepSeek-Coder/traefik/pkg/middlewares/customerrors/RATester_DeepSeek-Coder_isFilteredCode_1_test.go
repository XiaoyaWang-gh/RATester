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

	cc := &codeCatcher{
		caughtFilteredCode: true,
	}

	if !cc.isFilteredCode() {
		t.Errorf("Expected isFilteredCode to return true, but got false")
	}

	cc.caughtFilteredCode = false

	if cc.isFilteredCode() {
		t.Errorf("Expected isFilteredCode to return false, but got true")
	}
}
