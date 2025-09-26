package migration

import (
	"fmt"
	"testing"
)

func TestForeignCol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	foreign := m.ForeignCol("colname", "foreigncol", "foreigntable")

	if foreign.ForeignColumn != "foreigncol" {
		t.Errorf("Expected ForeignColumn to be 'foreigncol', but got %s", foreign.ForeignColumn)
	}

	if foreign.ForeignTable != "foreigntable" {
		t.Errorf("Expected ForeignTable to be 'foreigntable', but got %s", foreign.ForeignTable)
	}
}
