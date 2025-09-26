package ratelimit

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestTake_1(t *testing.T) {
	type args struct {
		amount uint
		ctx    *context.Context
	}
	tests := []struct {
		name string
		args args
		want bool
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

			l := &limiter{}
			if got := l.take(tt.args.amount, tt.args.ctx); got != tt.want {
				t.Errorf("limiter.take() = %v, want %v", got, tt.want)
			}
		})
	}
}
