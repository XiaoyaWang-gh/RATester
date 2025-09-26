package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestUpdateWithCtx_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		values orm.Params
	}
	tests := []struct {
		name    string
		d       *DoNothingQuerySetter
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

			got, err := tt.d.UpdateWithCtx(tt.args.ctx, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
