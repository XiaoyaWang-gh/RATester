package migration

import (
	"fmt"
	"testing"
)

func TestSetPrimary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	column := &Column{Name: "test"}
	migration := &Migration{}
	column.SetPrimary(migration)

	if len(migration.Primary) != 1 {
		t.Errorf("Expected 1 primary column, got %d", len(migration.Primary))
	}

	if migration.Primary[0] != column {
		t.Errorf("Expected primary column to be %v, got %v", column, migration.Primary[0])
	}
}
