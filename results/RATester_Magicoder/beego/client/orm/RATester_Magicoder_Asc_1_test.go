package orm

import (
	"fmt"
	"testing"
)

func TestAsc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Asc()

	if len(qb.tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(qb.tokens))
	}

	if qb.tokens[0] != "ASC" {
		t.Errorf("Expected 'ASC', got '%s'", qb.tokens[0])
	}
}
