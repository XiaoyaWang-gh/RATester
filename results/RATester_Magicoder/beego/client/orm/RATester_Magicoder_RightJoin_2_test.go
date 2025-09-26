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
	expected := fmt.Sprintf("RIGHT JOIN %s", quote+table+quote)
	qb.RightJoin(table)
	if qb.tokens[len(qb.tokens)-2] != "RIGHT JOIN" {
		t.Errorf("Expected 'RIGHT JOIN', got '%s'", qb.tokens[len(qb.tokens)-2])
	}
	if qb.tokens[len(qb.tokens)-1] != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.tokens[len(qb.tokens)-1])
	}
}
