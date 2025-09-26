package navigation

import (
	"fmt"
	"testing"
)

func TestGet_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &menuCache{}
	key := "key"
	apply := func(m Menu) {
		m.Add(&MenuEntry{})
	}
	menuLists := []Menu{}
	m, ok := c.get(key, apply, menuLists...)
	if !ok {
		t.Errorf("expected ok to be true")
	}
	if len(m) != 1 {
		t.Errorf("expected len(m) to be 1")
	}
}
