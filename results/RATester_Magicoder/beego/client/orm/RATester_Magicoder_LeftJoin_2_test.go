package orm

import (
	"fmt"
	"testing"
)

func TestLeftJoin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.LeftJoin("table1")

	expected := "LEFT JOIN table1"
	actual := qb.String()

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
