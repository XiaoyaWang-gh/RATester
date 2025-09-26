package orm

import (
	"context"
	"fmt"
	"testing"
)

func TestgetTxNameFromCtx_1(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Context_0",
			args: args{
				ctx: context.WithValue(context.Background(), TxNameKey, "test_tx_name"),
			},
			want: "test_tx_name",
		},
		{
			name: "Context_1",
			args: args{
				ctx: context.WithValue(context.Background(), TxNameKey, "another_tx_name"),
			},
			want: "another_tx_name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getTxNameFromCtx(tt.args.ctx); got != tt.want {
				t.Errorf("getTxNameFromCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
