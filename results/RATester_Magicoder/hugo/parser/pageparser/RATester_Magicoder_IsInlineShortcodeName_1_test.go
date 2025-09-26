package pageparser

import (
	"fmt"
	"testing"
)

func TestIsInlineShortcodeName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := Item{
		Type: tScNameInline,
	}

	if !item.IsInlineShortcodeName() {
		t.Errorf("Expected IsInlineShortcodeName to return true, but got false")
	}
}
