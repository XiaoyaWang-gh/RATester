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
	sub := "sub"
	alias := "alias"
	got := qb.Subquery(sub, alias)
	want := "(sub) AS alias"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
