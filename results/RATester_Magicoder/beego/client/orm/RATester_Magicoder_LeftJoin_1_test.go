package orm

import (
	"fmt"
	"testing"
)

func TestLeftJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	table := "test_table"
	expected := fmt.Sprintf("LEFT JOIN %s", quote+table+quote)
	qb.LeftJoin(table)
	if qb.tokens[len(qb.tokens)-2] != "LEFT JOIN" {
		t.Errorf("Expected LEFT JOIN, got %s", qb.tokens[len(qb.tokens)-2])
	}
	if qb.tokens[len(qb.tokens)-1] != expected {
		t.Errorf("Expected %s, got %s", expected, qb.tokens[len(qb.tokens)-1])
	}
}
