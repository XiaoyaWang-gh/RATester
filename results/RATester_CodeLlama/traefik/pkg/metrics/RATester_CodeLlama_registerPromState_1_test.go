package metrics

import (
	"context"
	"fmt"
	"testing"
)

func TestRegisterPromState_1(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_case_1",
			args: args{
				ctx: ctx,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := registerPromState(tt.args.ctx); got != tt.want {
				t.Errorf("registerPromState() = %v, want %v", got, tt.want)
			}
		})
	}
}
