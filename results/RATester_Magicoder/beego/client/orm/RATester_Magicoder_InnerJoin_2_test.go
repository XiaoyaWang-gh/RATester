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
	qb.InnerJoin("table")

	if len(qb.tokens) != 2 || qb.tokens[0] != "INNER JOIN" || qb.tokens[1] != "table" {
		t.Errorf("InnerJoin() failed. Expected: [INNER JOIN table], Actual: %v", qb.tokens)
	}
}
