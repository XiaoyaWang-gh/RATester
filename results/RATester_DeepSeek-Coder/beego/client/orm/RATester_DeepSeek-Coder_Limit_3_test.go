package orm

import (
	"fmt"
	"testing"
)

func TestLimit_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	limit := 10
	expected := "LIMIT 10"

	qb.Limit(limit)

	if len(qb.tokens) != 2 {
		t.Errorf("Expected length of tokens to be 2, got %d", len(qb.tokens))
	}

	if qb.tokens[1] != expected {
		t.Errorf("Expected token to be %s, got %s", expected, qb.tokens[1])
	}
}
