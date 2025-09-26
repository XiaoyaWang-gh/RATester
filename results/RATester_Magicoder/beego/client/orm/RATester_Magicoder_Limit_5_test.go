package orm

import (
	"fmt"
	"testing"
)

func TestLimit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Limit(10)

	if len(qb.tokens) != 2 {
		t.Errorf("Expected 2 tokens, got %d", len(qb.tokens))
	}

	if qb.tokens[0] != "LIMIT" {
		t.Errorf("Expected 'LIMIT', got '%s'", qb.tokens[0])
	}

	if qb.tokens[1] != "10" {
		t.Errorf("Expected '10', got '%s'", qb.tokens[1])
	}
}
