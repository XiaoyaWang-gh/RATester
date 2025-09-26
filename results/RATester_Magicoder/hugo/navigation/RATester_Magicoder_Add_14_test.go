package navigation

import (
	"fmt"
	"testing"
)

func TestAdd_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	menu := Menu{}
	entry := &MenuEntry{}
	menu = menu.Add(entry)

	if len(menu) != 1 {
		t.Errorf("Expected menu length to be 1, but got %d", len(menu))
	}

	if menu[0] != entry {
		t.Errorf("Expected menu entry to be added, but it was not")
	}
}
