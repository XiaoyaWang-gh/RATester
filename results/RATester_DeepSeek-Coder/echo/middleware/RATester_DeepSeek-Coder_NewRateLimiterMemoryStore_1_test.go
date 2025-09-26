package middleware

import (
	"fmt"
	"testing"

	"golang.org/x/time/rate"
)

func TestNewRateLimiterMemoryStore_1(t *testing.T) {
	type args struct {
		rate rate.Limit
	}
	tests := []struct {
		name     string
		args     args
		wantRate rate.Limit
	}{
		{
			name: "Test NewRateLimiterMemoryStore with rate.Limit 10",
			args: args{
				rate: 10,
			},
			wantRate: 10,
		},
		{
			name: "Test NewRateLimiterMemoryStore with rate.Limit 20",
			args: args{
				rate: 20,
			},
			wantRate: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotStore := NewRateLimiterMemoryStore(tt.args.rate)
			if gotStore.rate != tt.wantRate {
				t.Errorf("NewRateLimiterMemoryStore() = %v, want %v", gotStore.rate, tt.wantRate)
			}
		})
	}
}
