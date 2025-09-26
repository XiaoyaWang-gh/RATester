package orm

import (
	"fmt"
	"testing"
)

func TestUpdate_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "UPDATE table1, table2"
	result := qb.Update("table1", "table2").String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
