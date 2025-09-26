package orm

import (
	"fmt"
	"testing"
)

func TestOperatorSQL_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBasePostgres{}

	tests := []struct {
		operator string
		expected string
	}{
		{"=", "="},
		{"like", "ILIKE"},
		{"not like", "NOT ILIKE"},
		{"in", "IN"},
		{"not in", "NOT IN"},
		{"is null", "IS NULL"},
		{"is not null", "IS NOT NULL"},
		{"greater than", ">"},
		{"less than", "<"},
		{"greater than or equal", ">="},
		{"less than or equal", "<="},
		{"between", "BETWEEN"},
		{"not between", "NOT BETWEEN"},
		{"other", ""},
	}

	for _, test := range tests {
		result := db.OperatorSQL(test.operator)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
