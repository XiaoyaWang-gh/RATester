package order_clause

import (
	"fmt"
	"testing"
)

func TestColumn_1(t *testing.T) {
	type args struct {
		column string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with valid column name",
			args: args{column: "test_column"},
			want: "test_column",
		},
		{
			name: "Test with column name containing dot",
			args: args{column: "test.column"},
			want: "test.column",
		},
		{
			name: "Test with column name containing expression separator",
			args: args{column: "test$column"},
			want: "test.column",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			order := &Order{}
			Column(tt.args.column)(order)
			if order.column != tt.want {
				t.Errorf("Column() = %v, want %v", order.column, tt.want)
			}
		})
	}
}
