package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/utils"
)

func TestLoadRelatedWithCtx_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		name string
		args []utils.KV
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

			got, err := tt.d.LoadRelatedWithCtx(tt.args.ctx, tt.args.md, tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingOrm.LoadRelatedWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingOrm.LoadRelatedWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
