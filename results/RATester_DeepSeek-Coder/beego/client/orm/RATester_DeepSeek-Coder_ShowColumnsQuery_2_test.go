package orm

import (
	"fmt"
	"testing"
)

func TestShowColumnsQuery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseTidb{}
	table := "test_table"
	expected := fmt.Sprintf("SELECT COLUMN_NAME, COLUMN_TYPE, IS_NULLABLE FROM information_schema.Columns WHERE table_schema = DATABASE() AND table_name = '%s'", table)
	result := db.ShowColumnsQuery(table)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
