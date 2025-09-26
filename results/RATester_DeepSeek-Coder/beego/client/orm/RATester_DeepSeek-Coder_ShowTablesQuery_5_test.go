package orm

import (
	"fmt"
	"testing"
)

func TestShowTablesQuery_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}
	expected := "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema')"
	result := db.ShowTablesQuery()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
