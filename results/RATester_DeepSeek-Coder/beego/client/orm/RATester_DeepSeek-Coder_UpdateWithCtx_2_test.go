package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdateWithCtx_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		cols []string
	}
	tests := []struct {
		name    string
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

			o := &ormBase{}
			got, err := o.UpdateWithCtx(tt.args.ctx, tt.args.md, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ormBase.UpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ormBase.UpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
