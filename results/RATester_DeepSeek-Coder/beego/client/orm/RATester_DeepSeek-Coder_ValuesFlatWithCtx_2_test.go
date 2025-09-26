package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestValuesFlatWithCtx_2(t *testing.T) {
	type args struct {
		ctx    context.Context
		result *ParamsList
		expr   string
	}
	tests := []struct {
		name    string
		o       querySet
		args    args
		want    int64
		wantErr bool
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

			got, err := tt.o.ValuesFlatWithCtx(tt.args.ctx, tt.args.result, tt.args.expr)
			if (err != nil) != tt.wantErr {
				t.Errorf("querySet.ValuesFlatWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("querySet.ValuesFlatWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
