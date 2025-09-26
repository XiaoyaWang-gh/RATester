package orm

import (
	"fmt"
	"strings"
	"testing"
)

func TestShowColumnsQuery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseOracle{}
	table := "test_table"
	expected := fmt.Sprintf("SELECT COLUMN_NAME FROM ALL_TAB_COLUMNS WHERE TABLE_NAME ='%s'", strings.ToUpper(table))
	result := db.ShowColumnsQuery(table)
	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}
}
