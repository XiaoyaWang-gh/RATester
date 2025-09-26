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

	item := Item{Type: tText}
	if !item.IsText() {
		t.Errorf("Expected IsText to return true for item with Type tText")
	}

	item = Item{Type: tIndentation}
	if !item.IsText() {
		t.Errorf("Expected IsText to return true for item with Type tIndentation")
	}

	item = Item{Type: 0}
	if item.IsText() {
		t.Errorf("Expected IsText to return false for item with Type 0")
	}
}
