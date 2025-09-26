package orm

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestDoTxWithCtxAndOpts_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
		task func(ctx context.Context, txOrm TxOrmer) error
	}
	tests := []struct {
		name    string
		f       *filterOrmDecorator
		args    args
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

			if err := tt.f.DoTxWithCtxAndOpts(tt.args.ctx, tt.args.opts, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.DoTxWithCtxAndOpts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
