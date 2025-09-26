package navigation

import (
	"fmt"
	"testing"
)

func TestGetP_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &menuCache{}
	key := "key"
	apply := func(m *Menu) {}
	m := Menu{}
	menuLists := []Menu{m}

	_, ok := c.getP(key, apply, menuLists...)
	if ok {
		t.Errorf("getP() = (_, true), want (_, false)")
	}
}
