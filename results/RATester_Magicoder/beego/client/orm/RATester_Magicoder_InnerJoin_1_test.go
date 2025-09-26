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

	if len(qb.tokens) == 0 || qb.tokens[len(qb.tokens)-1] != expected {
		t.Errorf("InnerJoin() failed, expected %s, got %s", expected, qb.tokens[len(qb.tokens)-1])
	}
}
