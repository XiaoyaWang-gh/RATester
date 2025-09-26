package migration

import (
	"fmt"
	"testing"
)

func TestAddColumns_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}

	col1 := &Column{Name: "col1"}
	col2 := &Column{Name: "col2"}

	m.AddColumns(col1, col2)

	if len(m.Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(m.Columns))
	}

	if m.Columns[0].Name != "col1" || m.Columns[1].Name != "col2" {
		t.Errorf("Expected columns to be 'col1' and 'col2', got %s and %s", m.Columns[0].Name, m.Columns[1].Name)
	}
}
