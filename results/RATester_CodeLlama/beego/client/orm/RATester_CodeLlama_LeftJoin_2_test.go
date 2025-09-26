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
	qb.LeftJoin("table")
	if qb.tokens[0] != "LEFT JOIN" {
		t.Errorf("qb.tokens[0] = %v, want %v", qb.tokens[0], "LEFT JOIN")
	}
	if qb.tokens[1] != "table" {
		t.Errorf("qb.tokens[1] = %v, want %v", qb.tokens[1], "table")
	}
}
