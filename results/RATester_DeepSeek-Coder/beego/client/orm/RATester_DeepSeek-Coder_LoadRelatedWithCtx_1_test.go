package orm

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/core/utils"
)

func TestLoadRelatedWithCtx_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		md   interface{}
		name string
		args []utils.KV
	}
	tests := []struct {
		name    string
		o       *ormBase
		args    args
		wantNum int64
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

			gotNum, err := tt.o.LoadRelatedWithCtx(tt.args.ctx, tt.args.md, tt.args.name, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadRelatedWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotNum != tt.wantNum {
				t.Errorf("LoadRelatedWithCtx() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
		})
	}
}
