package middleware

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func TestNewRateLimiterMemoryStoreWithConfig_1(t *testing.T) {
	type args struct {
		config RateLimiterMemoryStoreConfig
	}
	tests := []struct {
		name          string
		args          args
		wantRate      rate.Limit
		wantBurst     int
		wantExpiresIn time.Duration
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

			gotStore := NewRateLimiterMemoryStoreWithConfig(tt.args.config)
			if gotStore.rate != tt.wantRate {
				t.Errorf("NewRateLimiterMemoryStoreWithConfig() = %v, want %v", gotStore.rate, tt.wantRate)
			}
			if gotStore.burst != tt.wantBurst {
				t.Errorf("NewRateLimiterMemoryStoreWithConfig() = %v, want %v", gotStore.burst, tt.wantBurst)
			}
			if gotStore.expiresIn != tt.wantExpiresIn {
				t.Errorf("NewRateLimiterMemoryStoreWithConfig() = %v, want %v", gotStore.expiresIn, tt.wantExpiresIn)
			}
		})
	}
}
