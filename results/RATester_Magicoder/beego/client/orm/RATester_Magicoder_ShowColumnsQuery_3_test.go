package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}
	table := "test_table"
	expected := fmt.Sprintf("SELECT column_name, data_type, is_nullable FROM information_schema.Columns where table_schema NOT IN ('pg_catalog', 'information_schema') and table_name = '%s'", table)
	result := db.ShowColumnsQuery(table)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
