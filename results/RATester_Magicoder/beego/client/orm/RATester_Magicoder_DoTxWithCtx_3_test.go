package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDoTxWithCtx_3(t *testing.T) {
	type args struct {
		ctx  context.Context
		task func(ctx context.Context, txOrm TxOrmer) error
	}
	tests := []struct {
		name    string
		o       *orm
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

			if err := tt.o.DoTxWithCtx(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("orm.DoTxWithCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
