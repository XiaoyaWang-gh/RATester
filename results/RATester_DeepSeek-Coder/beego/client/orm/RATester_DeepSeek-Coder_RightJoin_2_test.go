package orm

import (
	"fmt"
	"testing"
)

func TestRightJoin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	table := "test_table"
	expected := fmt.Sprintf("%s%s%s", quote, table, quote)

	qb.RightJoin(table)

	if len(qb.tokens) != 2 {
		t.Errorf("Expected length of tokens to be 2, got %d", len(qb.tokens))
	}

	if qb.tokens[0] != "RIGHT JOIN" {
		t.Errorf("Expected first token to be 'RIGHT JOIN', got %s", qb.tokens[0])
	}

	if qb.tokens[1] != expected {
		t.Errorf("Expected second token to be %s, got %s", expected, qb.tokens[1])
	}
}
