package orm

import (
	"fmt"
	"strconv"
	"testing"
)

func TestLimit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	limit := 10
	expected := "LIMIT 10"

	qb.Limit(limit)

	if qb.tokens[len(qb.tokens)-2] != "LIMIT" {
		t.Errorf("Expected 'LIMIT', got '%s'", qb.tokens[len(qb.tokens)-2])
	}

	if qb.tokens[len(qb.tokens)-1] != strconv.Itoa(limit) {
		t.Errorf("Expected '%s', got '%s'", expected, qb.tokens[len(qb.tokens)-1])
	}
}
