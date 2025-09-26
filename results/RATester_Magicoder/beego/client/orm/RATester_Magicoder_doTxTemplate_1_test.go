package orm

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestdoTxTemplate_1(t *testing.T) {
	type args struct {
		ctx  context.Context
		o    TxBeginner
		opts *sql.TxOptions
		task func(ctx context.Context, txOrm TxOrmer) error
	}
	tests := []struct {
		name    string
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

			if err := doTxTemplate(tt.args.ctx, tt.args.o, tt.args.opts, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("doTxTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
