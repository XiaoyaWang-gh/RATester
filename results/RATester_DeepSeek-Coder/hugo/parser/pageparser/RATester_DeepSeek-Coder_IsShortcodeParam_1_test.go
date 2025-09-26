package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortcodeParam_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := Item{
		Type: tScParam,
	}

	result := item.IsShortcodeParam()

	if !result {
		t.Errorf("Expected true, got false")
	}
}
