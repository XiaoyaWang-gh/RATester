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
	sub := "select * from users"
	alias := "u"
	expected := "(select * from users) AS u"
	actual := qb.Subquery(sub, alias)
	if actual != expected {
		t.Errorf("expected %s, actual %s", expected, actual)
	}
}
