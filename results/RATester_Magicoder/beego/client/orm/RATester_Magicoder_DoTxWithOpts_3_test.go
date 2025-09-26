package orm

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestDoTxWithOpts_3(t *testing.T) {
	type args struct {
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

			if err := tt.f.DoTxWithOpts(tt.args.opts, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.DoTxWithOpts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
