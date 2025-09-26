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
	columns := []*Column{}
	unique.AddColumnsToUnique(columns...)
	if len(unique.Columns) != 0 {
		t.Errorf("Expected 0 columns, got %d", len(unique.Columns))
	}
}
