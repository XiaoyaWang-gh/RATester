package pageparser

import (
	"fmt"
	"testing"
)

func TestIsShortcodeClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := Item{
		Type: tScClose,
	}

	result := item.IsShortcodeClose()

	if !result {
		t.Errorf("Expected IsShortcodeClose to return true, but got false")
	}
}
