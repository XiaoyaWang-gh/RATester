package navigation

import (
	"fmt"
	"testing"
)

func TestIsMenuCurrent_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m nopPageMenus
	var me MenuEntry

	// Test case 1: menuID is empty
	if m.IsMenuCurrent("", &me) {
		t.Errorf("Expected false, got true")
	}

	// Test case 2: inme is nil
	if m.IsMenuCurrent("menu1", nil) {
		t.Errorf("Expected false, got true")
	}

	// Test case 3: menuID is not empty and inme is not nil
	me = MenuEntry{
		MenuConfig: MenuConfig{
			Identifier: "menu1",
		},
	}
	if !m.IsMenuCurrent("menu1", &me) {
		t.Errorf("Expected true, got false")
	}
}
