package orm

import (
	"fmt"
	"testing"
)

func TestFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.From("table1", "table2")

	expected := "FROM table1, table2"
	if qb.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.String())
	}
}
