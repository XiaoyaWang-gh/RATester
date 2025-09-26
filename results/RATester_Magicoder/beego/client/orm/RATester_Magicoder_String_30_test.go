package orm

import (
	"fmt"
	"testing"
)

func TestString_30(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{
		tokens: []string{"SELECT", "*", "FROM", "table"},
	}

	expected := "SELECT * FROM table"
	result := qb.String()

	if result != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, result)
	}

	if len(qb.tokens) != 0 {
		t.Errorf("Expected tokens to be empty, but got %v", qb.tokens)
	}
}
