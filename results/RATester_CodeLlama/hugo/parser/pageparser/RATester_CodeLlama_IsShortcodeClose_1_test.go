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

	i := Item{Type: tScClose}
	if !i.IsShortcodeClose() {
		t.Errorf("IsShortcodeClose failed")
	}
}
