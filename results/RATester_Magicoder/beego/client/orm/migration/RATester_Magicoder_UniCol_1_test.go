package migration

import (
	"fmt"
	"testing"
)

func TestUniCol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	uni := "uni"
	name := "name"
	col := m.UniCol(uni, name)

	if col.Name != name {
		t.Errorf("Expected name to be %s, but got %s", name, col.Name)
	}

	unique := &Unique{}
	for _, u := range m.Uniques {
		if u.Definition == uni {
			unique = u
		}
	}

	if unique.Definition != uni {
		t.Errorf("Expected unique definition to be %s, but got %s", uni, unique.Definition)
	}

	if len(unique.Columns) != 1 {
		t.Errorf("Expected unique columns length to be 1, but got %d", len(unique.Columns))
	}

	if unique.Columns[0].Name != name {
		t.Errorf("Expected unique column name to be %s, but got %s", name, unique.Columns[0].Name)
	}
}
