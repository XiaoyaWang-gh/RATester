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
	sub := "SELECT * FROM users"
	alias := "u"
	expected := "(SELECT * FROM users) AS u"
	result := qb.Subquery(sub, alias)
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
