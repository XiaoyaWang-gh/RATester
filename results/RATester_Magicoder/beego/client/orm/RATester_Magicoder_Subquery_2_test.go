package orm

import (
	"fmt"
	"testing"
)

func TestSubquery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	sub := "SELECT * FROM table"
	alias := "alias"
	expected := fmt.Sprintf("(%s) AS %s", sub, alias)
	result := qb.Subquery(sub, alias)
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
