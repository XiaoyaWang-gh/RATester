package mock

import (
	"context"
	"fmt"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

func TestValuesListWithCtx_1(t *testing.T) {
	type args struct {
		ctx     context.Context
		results *[]orm.ParamsList
		exprs   []string
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

			got, err := tt.d.ValuesListWithCtx(tt.args.ctx, tt.args.results, tt.args.exprs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DoNothingQuerySetter.ValuesListWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DoNothingQuerySetter.ValuesListWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
