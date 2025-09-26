package orm

import (
	"fmt"
	"testing"
)

func TestDelete_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Delete("table1", "table2")
	expected := "DELETE FROM table1, table2"
	result := qb.String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
