package navigation

import (
	"fmt"
	"testing"
)

func TestMenus_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	nop := nopPageMenus(0)
	menus := nop.Menus()
	if menus != nil {
		t.Errorf("Expected nil, got %v", menus)
	}
}
