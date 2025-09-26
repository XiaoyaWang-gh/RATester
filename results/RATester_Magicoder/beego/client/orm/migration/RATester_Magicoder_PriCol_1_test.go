package migration

import (
	"fmt"
	"testing"
)

func TestPriCol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	col := m.PriCol("test_column")

	if col.Name != "test_column" {
		t.Errorf("Expected column name to be 'test_column', but got '%s'", col.Name)
	}

	if len(m.Primary) != 1 {
		t.Errorf("Expected 1 primary column, but got %d", len(m.Primary))
	}

	if m.Primary[0] != col {
		t.Errorf("Expected primary column to be the one returned by PriCol")
	}
}
