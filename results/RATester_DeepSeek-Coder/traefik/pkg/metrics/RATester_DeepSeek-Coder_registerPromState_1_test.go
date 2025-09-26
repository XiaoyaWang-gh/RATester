package metrics

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestRegisterPromState_1(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx = logger.WithContext(ctx)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test registerPromState with context.Background()",
			args: args{ctx: ctx},
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
