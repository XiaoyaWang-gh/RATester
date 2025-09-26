package orm

import (
	"fmt"
	"testing"
)

func TestSubquery_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	subquery := "SELECT * FROM table"
	alias := "alias"
	expected := fmt.Sprintf("(%s) AS %s", subquery, alias)
	result := qb.Subquery(subquery, alias)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
