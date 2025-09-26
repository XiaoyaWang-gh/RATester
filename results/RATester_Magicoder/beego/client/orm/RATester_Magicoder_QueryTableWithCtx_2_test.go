package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestQueryTableWithCtx_2(t *testing.T) {
	ctx := context.Background()
	orm := &DoNothingOrm{}

	tests := []struct {
		name string
		ctx  context.Context
		want QuerySeter
	}{
		{
			name: "Test case 1",
			ctx:  ctx,
			want: orm.QueryTableWithCtx(ctx, "test_table"),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := orm.QueryTableWithCtx(tt.ctx, "test_table")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryTableWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
