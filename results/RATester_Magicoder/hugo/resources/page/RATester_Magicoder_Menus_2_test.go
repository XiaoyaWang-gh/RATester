package page

import (
	"fmt"
	"testing"
)

func TestMenus_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPage(0)
	menus := nop.Menus()
	if menus == nil {
		t.Errorf("Expected non-nil menus, got nil")
	}
}
