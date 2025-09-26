package migration

import (
	"fmt"
	"testing"
)

func TestAddPrimary_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	primary := &Column{Name: "id"}
	m.AddPrimary(primary)

	if len(m.Primary) != 1 {
		t.Errorf("Expected 1 primary column, got %d", len(m.Primary))
	}

	if m.Primary[0].Name != "id" {
		t.Errorf("Expected primary column name to be 'id', got '%s'", m.Primary[0].Name)
	}
}
