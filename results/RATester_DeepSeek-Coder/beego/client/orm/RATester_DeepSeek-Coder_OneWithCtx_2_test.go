package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestOneWithCtx_2(t *testing.T) {
	type args struct {
		ctx       context.Context
		container interface{}
		cols      []string
	}
	tests := []struct {
		name    string
		o       querySet
		args    args
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

			if err := tt.o.OneWithCtx(tt.args.ctx, tt.args.container, tt.args.cols...); (err != nil) != tt.wantErr {
				t.Errorf("querySet.OneWithCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
