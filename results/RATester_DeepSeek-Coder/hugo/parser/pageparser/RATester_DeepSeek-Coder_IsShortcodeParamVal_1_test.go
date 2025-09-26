package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortcodeParamVal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := Item{
		Type: tScParamVal,
	}

	result := item.IsShortcodeParamVal()

	if !result {
		t.Errorf("Expected true, got false")
	}
}
