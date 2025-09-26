package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestBeginWithCtxAndOpts_2(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
	}
	tests := []struct {
		name    string
		f       *filterOrmDecorator
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

			got, err := tt.f.BeginWithCtxAndOpts(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.BeginWithCtxAndOpts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterOrmDecorator.BeginWithCtxAndOpts() = %v, want %v", got, tt.want)
			}
		})
	}
}
