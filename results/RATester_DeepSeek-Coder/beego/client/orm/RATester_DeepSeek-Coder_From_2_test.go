package orm

import (
	"fmt"
	"testing"
)

func TestFrom_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	expected := "FROM table1, table2"
	result := qb.From("table1", "table2").String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
