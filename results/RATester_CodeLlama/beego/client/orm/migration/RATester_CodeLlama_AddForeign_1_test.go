package migration

import (
	"fmt"
	"testing"
)

func TestAddForeign_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	foreign := &Foreign{}
	m.AddForeign(foreign)
	if len(m.Foreigns) != 1 {
		t.Errorf("AddForeign failed")
	}
}
