package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestUpdateWithCtx_5(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		cols []string
	}
	tests := []struct {
		name    string
		d       *DoNothingOrm
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

			got, err := tt.d.UpdateWithCtx(tt.args.ctx, tt.args.md, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingOrm.UpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingOrm.UpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
