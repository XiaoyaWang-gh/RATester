package navigation

import (
	"fmt"
	"testing"
)

func TestIsEqual_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &MenuEntry{
		MenuConfig: MenuConfig{
			Identifier: "id1",
			Parent:     "parent1",
		},
	}
	inme := &MenuEntry{
		MenuConfig: MenuConfig{
			Identifier: "id1",
			Parent:     "parent1",
		},
	}
	if !m.isEqual(inme) {
		t.Errorf("m.isEqual(inme) = false, want true")
	}
}
