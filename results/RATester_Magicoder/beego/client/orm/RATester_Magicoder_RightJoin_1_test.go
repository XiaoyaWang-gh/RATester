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
	qb.RightJoin("table")

	if len(qb.tokens) != 2 || qb.tokens[0] != "RIGHT JOIN" || qb.tokens[1] != "table" {
		t.Errorf("RightJoin() failed. Expected tokens: [\"RIGHT JOIN\", \"table\"], got: %v", qb.tokens)
	}
}
