package retry

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestRetry_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		initial       time.Duration
		maxBackoff    time.Duration
		multiplier    float64
		maxRetryCount int
		f             func() error
		wantErr       error
	}{
		{
			name:          "success on first try",
			initial:       10 * time.Millisecond,
			maxBackoff:    1 * time.Second,
			multiplier:    1.5,
			maxRetryCount: 3,
			f: func() error {
				return nil
			},
			wantErr: nil,
		},
		{
			name:          "fail after retries",
			initial:       10 * time.Millisecond,
			maxBackoff:    1 * time.Second,
			multiplier:    1.5,
			maxRetryCount: 3,
			f: func() error {
				return errors.New("always failing")
			},
			wantErr: errors.New("always failing"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := &ExponentialBackoff{
				InitialInterval: tt.initial,
				MaxBackoffTime:  tt.maxBackoff,
				Multiplier:      tt.multiplier,
				MaxRetryCount:   tt.maxRetryCount,
			}
			err := e.Retry(tt.f)
			if err != tt.wantErr {
				t.Errorf("Retry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
