package order_clause

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseOrder_1(t *testing.T) {
	tests := []struct {
		name        string
		expressions []string
		want        []*Order
	}{
		{
			name:        "test case 1",
			expressions: []string{"column1", "-column2"},
			want: []*Order{
				{
					column: "column1",
					sort:   Ascending,
				},
				{
					column: "column2",
					sort:   Descending,
				},
			},
		},
		{
			name:        "test case 2",
			expressions: []string{"column3", "column4"},
			want: []*Order{
				{
					column: "column3",
					sort:   Ascending,
				},
				{
					column: "column4",
					sort:   Ascending,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := ParseOrder(tt.expressions...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
