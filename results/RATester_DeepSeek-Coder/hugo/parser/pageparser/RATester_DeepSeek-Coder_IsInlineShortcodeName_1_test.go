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

	result := item.IsInlineShortcodeName()

	if !result {
		t.Errorf("Expected true, got false")
	}
}
