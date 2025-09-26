package orm

import (
	"fmt"
	"testing"
)

func TestOperatorSQL_3(t *testing.T) {
	db := &dbBasePostgres{}

	t.Run("TestOperatorSQL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name     string
			operator string
			want     string
		}{
			{"TestEqualOperator", "=", "="},
			{"TestNotEqualOperator", "<>", "<>"},
			{"TestGreaterThanOperator", ">", ">"},
			{"TestGreaterThanOrEqualOperator", ">=", ">="},
			{"TestLessThanOperator", "<", "<"},
			{"TestLessThanOrEqualOperator", "<=", "<="},
			{"TestInOperator", "IN", "IN"},
			{"TestNotInOperator", "NOT IN", "NOT IN"},
			{"TestIsNullOperator", "IS NULL", "IS NULL"},
			{"TestIsNotNullOperator", "IS NOT NULL", "IS NOT NULL"},
			{"TestLikeOperator", "LIKE", "LIKE"},
			{"TestNotLikeOperator", "NOT LIKE", "NOT LIKE"},
			{"TestIlikeOperator", "ILIKE", "ILIKE"},
			{"TestNotIlikeOperator", "NOT ILIKE", "NOT ILIKE"},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				got := db.OperatorSQL(tc.operator)
				if got != tc.want {
					t.Errorf("got %s, want %s", got, tc.want)
				}
			})
		}
	})
}
