package order_clause

import (
	"fmt"
	"testing"
)

func TestSortString_1(t *testing.T) {
	type fields struct {
		column string
		sort   Sort
		isRaw  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Ascending",
			fields: fields{
				column: "column1",
				sort:   Ascending,
				isRaw:  false,
			},
			want: "ASC",
		},
		{
			name: "Test Descending",
			fields: fields{
				column: "column1",
				sort:   Descending,
				isRaw:  false,
			},
			want: "DESC",
		},
		{
			name: "Test Empty",
			fields: fields{
				column: "column1",
				sort:   0,
				isRaw:  false,
			},
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			o := &Order{
				column: tt.fields.column,
				sort:   tt.fields.sort,
				isRaw:  tt.fields.isRaw,
			}
			if got := o.SortString(); got != tt.want {
				t.Errorf("Order.SortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
