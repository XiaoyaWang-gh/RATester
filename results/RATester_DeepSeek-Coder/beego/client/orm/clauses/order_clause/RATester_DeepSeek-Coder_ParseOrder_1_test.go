package order_clause

import (
	"fmt"
	"testing"
)

func TestParseOrder_1(t *testing.T) {
	type args struct {
		expressions []string
	}
	tests := []struct {
		name string
		args args
		want []*Order
	}{
		{
			name: "Test case 1",
			args: args{
				expressions: []string{"column1", "-column2"},
			},
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
			name: "Test case 2",
			args: args{
				expressions: []string{"column1", "column2"},
			},
			want: []*Order{
				{
					column: "column1",
					sort:   Ascending,
				},
				{
					column: "column2",
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

			got := ParseOrder(tt.args.expressions...)
			for i, order := range got {
				if order.GetColumn() != tt.want[i].GetColumn() || order.GetSort() != tt.want[i].GetSort() {
					t.Errorf("ParseOrder() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
