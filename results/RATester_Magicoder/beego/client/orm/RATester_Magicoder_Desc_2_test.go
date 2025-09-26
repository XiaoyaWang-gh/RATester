package orm

import (
	"fmt"
	"testing"
)

func TestDesc_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Desc()

	if len(qb.tokens) != 1 {
		t.Errorf("Expected 1 token, got %d", len(qb.tokens))
	}

	if qb.tokens[0] != "DESC" {
		t.Errorf("Expected 'DESC', got '%s'", qb.tokens[0])
	}
}
