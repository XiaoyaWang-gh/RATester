package orm

import (
	"fmt"
	"testing"
)

func TestRightJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	table := "test_table"
	qb.RightJoin(table)

	expected := "RIGHT JOIN test_table"
	actual := qb.String()

	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
