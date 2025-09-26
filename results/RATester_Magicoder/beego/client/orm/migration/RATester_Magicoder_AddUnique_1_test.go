package migration

import (
	"fmt"
	"testing"
)

func TestAddUnique_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Migration instance
	m := &Migration{}

	// Create a new Unique instance
	unique := &Unique{
		Definition: "unique_definition",
		Columns: []*Column{
			{Name: "column1"},
			{Name: "column2"},
		},
	}

	// Add the unique to the migration
	m.AddUnique(unique)

	// Assert that the unique was added to the migration
	if len(m.Uniques) != 1 {
		t.Errorf("Expected 1 unique, got %d", len(m.Uniques))
	}

	// Assert that the unique has the correct definition and columns
	if m.Uniques[0].Definition != "unique_definition" {
		t.Errorf("Expected unique definition 'unique_definition', got '%s'", m.Uniques[0].Definition)
	}
	if len(m.Uniques[0].Columns) != 2 {
		t.Errorf("Expected 2 columns, got %d", len(m.Uniques[0].Columns))
	}
	if m.Uniques[0].Columns[0].Name != "column1" {
		t.Errorf("Expected column1, got '%s'", m.Uniques[0].Columns[0].Name)
	}
	if m.Uniques[0].Columns[1].Name != "column2" {
		t.Errorf("Expected column2, got '%s'", m.Uniques[0].Columns[1].Name)
	}
}
