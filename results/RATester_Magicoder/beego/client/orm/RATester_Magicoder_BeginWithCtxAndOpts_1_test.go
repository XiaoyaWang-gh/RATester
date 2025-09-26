package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestBeginWithCtxAndOpts_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
	}
	tests := []struct {
		name    string
		o       *orm
		args    args
		want    TxOrmer
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

			got, err := tt.o.BeginWithCtxAndOpts(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("orm.BeginWithCtxAndOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("orm.BeginWithCtxAndOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
