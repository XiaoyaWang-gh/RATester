package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestQueryTableWithCtx_2(t *testing.T) {
	type args struct {
		ctx                  context.Context
		ptrStructOrTableName interface{}
	}
	tests := []struct {
		name string
		args args
		want QuerySeter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &DoNothingOrm{}
			if got := d.QueryTableWithCtx(tt.args.ctx, tt.args.ptrStructOrTableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryTableWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
