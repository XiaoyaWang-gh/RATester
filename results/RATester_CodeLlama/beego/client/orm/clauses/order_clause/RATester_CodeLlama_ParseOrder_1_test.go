package order_clause

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseOrder_1(t *testing.T) {
	var tests = []struct {
		name        string
		expressions []string
		want        []*Order
	}{
		{
			name:        "test1",
			expressions: []string{"id", "name"},
			want:        []*Order{{column: "id", sort: Ascending}, {column: "name", sort: Ascending}},
		},
		{
			name:        "test2",
			expressions: []string{"-id", "name"},
			want:        []*Order{{column: "id", sort: Descending}, {column: "name", sort: Ascending}},
		},
		{
			name:        "test3",
			expressions: []string{"id", "-name"},
			want:        []*Order{{column: "id", sort: Ascending}, {column: "name", sort: Descending}},
		},
		{
			name:        "test4",
			expressions: []string{"-id", "-name"},
			want:        []*Order{{column: "id", sort: Descending}, {column: "name", sort: Descending}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ParseOrder(tt.expressions...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
