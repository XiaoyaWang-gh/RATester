package pageparser

import (
	"fmt"
	"testing"
)

func TestIsIndentation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := Item{
		Type: tIndentation,
	}

	if !item.IsIndentation() {
		t.Errorf("Expected IsIndentation to return true, but got false")
	}
}
