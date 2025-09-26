package migration

import (
	"fmt"
	"testing"
)

func TestNewCol_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	colName := "test_column"
	col := m.NewCol(colName)

	if col.Name != colName {
		t.Errorf("Expected column name to be %s, but got %s", colName, col.Name)
	}

	if len(m.Columns) != 1 {
		t.Errorf("Expected migration to have 1 column, but got %d", len(m.Columns))
	}

	if m.Columns[0] != col {
		t.Errorf("Expected first column in migration to be the newly created column")
	}
}
