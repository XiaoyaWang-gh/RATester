package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDeleteWithCtx_3(t *testing.T) {
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

			got, err := tt.d.DeleteWithCtx(tt.args.ctx, tt.args.md, tt.args.cols...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingOrm.DeleteWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingOrm.DeleteWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
