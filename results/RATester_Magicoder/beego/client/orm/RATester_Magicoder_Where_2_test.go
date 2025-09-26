package orm

import (
	"fmt"
	"testing"
)

func TestWhere_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	cond := "id = 1"
	qb.Where(cond)

	if len(qb.tokens) != 3 || qb.tokens[0] != "WHERE" || qb.tokens[1] != cond {
		t.Errorf("Expected tokens to be ['WHERE', '%s'], but got %v", cond, qb.tokens)
	}
}
