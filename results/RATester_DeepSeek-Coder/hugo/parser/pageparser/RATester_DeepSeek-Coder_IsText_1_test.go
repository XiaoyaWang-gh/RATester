package pageparser

import (
	"fmt"
	"testing"
)

func TestIsText_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := Item{
		Type: tText,
	}

	if !item.IsText() {
		t.Errorf("Expected IsText to return true for tText")
	}

	item.Type = tIndentation

	if !item.IsText() {
		t.Errorf("Expected IsText to return true for tIndentation")
	}

	item.Type = 0

	if item.IsText() {
		t.Errorf("Expected IsText to return false for 0")
	}
}
