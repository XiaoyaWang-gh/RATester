package order_clause

import (
	"fmt"
	"reflect"
	"testing"
)

func TestColumn_1(t *testing.T) {
	type args struct {
		column string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		{
			name: "TestColumn",
			args: args{
				column: "test_column",
			},
			want: func(order *Order) {
				order.column = "test_column"
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

			if got := Column(tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Column() = %v, want %v", got, tt.want)
			}
		})
	}
}
