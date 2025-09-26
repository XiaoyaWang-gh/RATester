package orm

import (
	"fmt"
	"testing"
)

func TestAnd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.And("condition")

	if len(qb.tokens) != 2 || qb.tokens[0] != "AND" || qb.tokens[1] != "condition" {
		t.Errorf("Expected tokens to be ['AND', 'condition'], but got %v", qb.tokens)
	}
}
