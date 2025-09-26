package orm

import (
	"fmt"
	"testing"
)

func TestInnerJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	table := "test_table"
	expected := fmt.Sprintf("INNER JOIN %s", quote+table+quote)

	qb.InnerJoin(table)

	if len(qb.tokens) != 2 {
		t.Errorf("Expected length of tokens to be 2, got %d", len(qb.tokens))
	}

	if qb.tokens[1] != expected {
		t.Errorf("Expected token to be %s, got %s", expected, qb.tokens[1])
	}
}
