package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestDoTxWithCtx_2(t *testing.T) {
	type args struct {
		ctx  context.Context
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

			if err := tt.f.DoTxWithCtx(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("filterOrmDecorator.DoTxWithCtx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
