package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesWithCtx_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		results *[]orm.Params
		exprs   []string
	}
	tests := []struct {
		name    string
		d       *DoNothingQuerySetter
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "test_values_with_ctx_0",
			d:    &DoNothingQuerySetter{},
			args: args{
				ctx:     context.Background(),
				results: &[]orm.Params{},
				exprs:   []string{},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := tt.d.ValuesWithCtx(tt.args.ctx, tt.args.results, tt.args.exprs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValuesWithCtx() got = %v, want %v", got, tt.want)
			}
		})
	}
}
