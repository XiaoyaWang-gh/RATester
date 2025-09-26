package orm

import (
	"fmt"
	"testing"
)

func TestOperatorSQL_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	db := &dbBaseMysql{}

	tests := []struct {
		operator string
		expected string
	}{
		{"=", "="},
		{"like", "LIKE"},
		{">", ">"},
		{"<", "<"},
		{"<>", "<>"},
		{"in", "IN"},
		{"not in", "NOT IN"},
		{"is null", "IS NULL"},
		{"is not null", "IS NOT NULL"},
		{"between", "BETWEEN"},
		{"not between", "NOT BETWEEN"},
		{"", ""},
	}

	for _, test := range tests {
		result := db.OperatorSQL(test.operator)
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
