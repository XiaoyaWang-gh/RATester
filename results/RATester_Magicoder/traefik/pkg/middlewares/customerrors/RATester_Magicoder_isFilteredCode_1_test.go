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
		caughtFilteredCode: false,
	}

	result := cc.isFilteredCode()
	if result != false {
		t.Errorf("Expected false, got %v", result)
	}

	cc.caughtFilteredCode = true
	result = cc.isFilteredCode()
	if result != true {
		t.Errorf("Expected true, got %v", result)
	}
}
