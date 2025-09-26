package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}
	table := "test_table"
	expected := fmt.Sprintf("pragma table_info('%s')", table)
	result := db.ShowColumnsQuery(table)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
