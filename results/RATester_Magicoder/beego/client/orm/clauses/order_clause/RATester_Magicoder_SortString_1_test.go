package order_clause

import (
	"fmt"
	"testing"
)

func TestSortString_1(t *testing.T) {
	tests := []struct {
		name     string
		order    *Order
		expected string
	}{
		{
			name: "Ascending",
			order: &Order{
				column: "column",
				sort:   Ascending,
				isRaw:  false,
			},
			expected: "ASC",
		},
		{
			name: "Descending",
			order: &Order{
				column: "column",
				sort:   Descending,
				isRaw:  false,
			},
			expected: "DESC",
		},
		{
			name: "Default",
			order: &Order{
				column: "column",
				sort:   0,
				isRaw:  false,
			},
			expected: ``,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.order.SortString(); got != tt.expected {
				t.Errorf("SortString() = %v, want %v", got, tt.expected)
			}
		})
	}
}
