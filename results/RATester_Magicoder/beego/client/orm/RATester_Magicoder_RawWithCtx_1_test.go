package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestRawWithCtx_1(t *testing.T) {
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want RawSeter
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

			o := &ormBase{}
			if got := o.RawWithCtx(tt.args.ctx, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
