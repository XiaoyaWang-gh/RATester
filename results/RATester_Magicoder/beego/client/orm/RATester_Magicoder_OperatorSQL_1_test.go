package orm

import (
	"fmt"
	"testing"
)

func TestOperatorSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseSqlite{}

	tests := []struct {
		operator string
		expected string
	}{
		{"=", "="},
		{"like", "LIKE"},
		{"in", "IN"},
		{"not in", "NOT IN"},
		{"is null", "IS NULL"},
		{"is not null", "IS NOT NULL"},
		{"not", "NOT"},
		{"and", "AND"},
		{"or", "OR"},
		{"", ""},
	}

	for _, test := range tests {
		result := db.OperatorSQL(test.operator)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
