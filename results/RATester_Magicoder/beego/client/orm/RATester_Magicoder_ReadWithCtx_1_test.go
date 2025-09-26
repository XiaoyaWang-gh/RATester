package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestReadWithCtx_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		cols []string
	}
	tests := []struct {
		name    string
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

			f := &filterOrmDecorator{}
			if err := f.ReadWithCtx(tt.args.ctx, tt.args.md, tt.args.cols...); (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.ReadWithCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
