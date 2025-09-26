package retry

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewExponentialBackoff_1(t *testing.T) {
	tests := []struct {
		name   string
		config []Config
		want   *ExponentialBackoff
	}{
		{
			name: "Test with default config",
			config: []Config{
				{},
			},
			want: &ExponentialBackoff{
				InitialInterval: 1 * time.Second,
				MaxBackoffTime:  32 * time.Second,
				Multiplier:      2.0,
				MaxRetryCount:   10,
				currentInterval: 1 * time.Second,
			},
		},
		{
			name: "Test with custom config",
			config: []Config{
				{
					InitialInterval: 2 * time.Second,
					MaxBackoffTime:  16 * time.Second,
					Multiplier:      3.0,
					MaxRetryCount:   5,
					currentInterval: 2 * time.Second,
				},
			},
			want: &ExponentialBackoff{
				InitialInterval: 2 * time.Second,
				MaxBackoffTime:  16 * time.Second,
				Multiplier:      3.0,
				MaxRetryCount:   5,
				currentInterval: 2 * time.Second,
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

			got := NewExponentialBackoff(tt.config...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExponentialBackoff() = %v, want %v", got, tt.want)
			}
		})
	}
}
