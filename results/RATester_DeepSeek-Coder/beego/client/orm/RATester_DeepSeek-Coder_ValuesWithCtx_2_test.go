package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestValuesWithCtx_2(t *testing.T) {
	type args struct {
		ctx     context.Context
		results *[]Params
		exprs   []string
	}
	tests := []struct {
		name    string
		o       querySet
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

			got, err := tt.o.ValuesWithCtx(tt.args.ctx, tt.args.results, tt.args.exprs...)
			if (err != nil) != tt.wantErr {
				t.Errorf("querySet.ValuesWithCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("querySet.ValuesWithCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
