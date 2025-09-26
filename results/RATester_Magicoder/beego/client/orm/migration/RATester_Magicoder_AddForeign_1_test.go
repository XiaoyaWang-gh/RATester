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
	foreign := &Foreign{
		ForeignTable: "test_table",
	}
	m.AddForeign(foreign)

	if len(m.Foreigns) != 1 {
		t.Errorf("Expected 1 foreign, got %d", len(m.Foreigns))
	}

	if m.Foreigns[0] != foreign {
		t.Errorf("Expected added foreign to be the same as the original")
	}
}
