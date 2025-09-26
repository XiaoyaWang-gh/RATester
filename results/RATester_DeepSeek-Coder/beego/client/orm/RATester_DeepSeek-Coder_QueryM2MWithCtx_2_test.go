package orm

import (
	"context"
	"fmt"
	"reflect"
	"testing"
)

func TestQueryM2MWithCtx_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		name string
	}
	tests := []struct {
		name string
		args args
		want QueryM2Mer
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
			if got := d.QueryM2MWithCtx(tt.args.ctx, tt.args.md, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DoNothingOrm.QueryM2MWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
