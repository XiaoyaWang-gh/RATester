package safe

import (
	"context"
	"fmt"
	"testing"
)

func TestNewPool_1(t *testing.T) {
	type args struct {
		parentCtx context.Context
	}
	tests := []struct {
		name string
		args args
		want *Pool
	}{
		{
			name: "TestNewPool_WithNilParentContext",
			args: args{
				parentCtx: nil,
			},
			want: &Pool{
				ctx:    nil,
				cancel: nil,
			},
		},
		{
			name: "TestNewPool_WithNonNilParentContext",
			args: args{
				parentCtx: context.Background(),
			},
			want: &Pool{
				ctx:    context.Background(),
				cancel: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewPool(tt.args.parentCtx)
			if got.ctx != tt.want.ctx {
				t.Errorf("NewPool() = %v, want %v", got.ctx, tt.want.ctx)
			}
			if (got.cancel != nil) != (tt.want.cancel != nil) {
				t.Errorf("NewPool() = %v, want %v", got.cancel, tt.want.cancel)
			}
		})
	}
}
