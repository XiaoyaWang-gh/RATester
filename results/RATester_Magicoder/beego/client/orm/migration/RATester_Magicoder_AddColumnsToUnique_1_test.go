package migration

import (
	"fmt"
	"testing"
)

func TestAddColumnsToUnique_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	unique := &Unique{}
	column1 := &Column{Name: "column1"}
	column2 := &Column{Name: "column2"}

	unique.AddColumnsToUnique(column1, column2)

	if len(unique.Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(unique.Columns))
	}

	if unique.Columns[0].Name != "column1" {
		t.Errorf("Expected first column name to be 'column1', got %s", unique.Columns[0].Name)
	}

	if unique.Columns[1].Name != "column2" {
		t.Errorf("Expected second column name to be 'column2', got %s", unique.Columns[1].Name)
	}
}
