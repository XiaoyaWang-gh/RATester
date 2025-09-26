package orm

import (
	"fmt"
	"testing"
)

func TestInnerJoin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	table := "test_table"
	qb.InnerJoin(table)

	expected := "INNER JOIN " + table
	actual := qb.tokens[len(qb.tokens)-2] + " " + qb.tokens[len(qb.tokens)-1]

	if expected != actual {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
