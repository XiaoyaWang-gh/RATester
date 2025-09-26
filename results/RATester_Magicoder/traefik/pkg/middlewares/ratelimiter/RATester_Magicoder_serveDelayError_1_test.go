package ratelimiter

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServeDelayError_1(t *testing.T) {
	type args struct {
		ctx   context.Context
		w     http.ResponseWriter
		delay time.Duration
	}
	tests := []struct {
		name string
		args args
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

			rl := &rateLimiter{}
			rl.serveDelayError(tt.args.ctx, tt.args.w, tt.args.delay)
		})
	}
}
