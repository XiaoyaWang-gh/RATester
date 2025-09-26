package orm

import (
	"fmt"
	"strconv"
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
	qb.Limit(limit)

	if len(qb.tokens) != 2 {
		t.Errorf("Expected tokens length to be 2, got %d", len(qb.tokens))
	}

	if qb.tokens[0] != "LIMIT" {
		t.Errorf("Expected first token to be 'LIMIT', got %s", qb.tokens[0])
	}

	parsedLimit, err := strconv.Atoi(qb.tokens[1])
	if err != nil {
		t.Errorf("Expected second token to be a valid integer, got %s", qb.tokens[1])
	}

	if parsedLimit != limit {
		t.Errorf("Expected limit to be %d, got %d", limit, parsedLimit)
	}
}
