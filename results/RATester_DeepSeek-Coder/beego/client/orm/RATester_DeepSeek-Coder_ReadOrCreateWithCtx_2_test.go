package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestReadOrCreateWithCtx_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		col1 string
		cols []string
	}
	tests := []struct {
		name    string
		f       *filterOrmDecorator
		args    args
		want    bool
		want1   int64
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

			got, got1, err := tt.f.ReadOrCreateWithCtx(tt.args.ctx, tt.args.md, tt.args.col1, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadOrCreateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadOrCreateWithCtx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ReadOrCreateWithCtx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
