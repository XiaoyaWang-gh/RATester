package migration

import (
	"fmt"
	"testing"
)

func TestExec_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{
		sqls: []string{"CREATE TABLE test_table (id INT, name VARCHAR(255))"},
	}

	err := m.Exec("test_migration", "pending")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
