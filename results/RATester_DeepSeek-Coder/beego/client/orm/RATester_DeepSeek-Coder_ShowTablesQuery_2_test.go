package orm

import (
	"fmt"
	"testing"
)

func TestShowTablesQuery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}
	expected := "SELECT name FROM sqlite_master WHERE type = 'table'"
	result := db.ShowTablesQuery()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
