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

	m := Menu{}
	me := &MenuEntry{}
	m.Add(me)
	if len(m) != 1 {
		t.Errorf("m.Add(me) failed")
	}
}
