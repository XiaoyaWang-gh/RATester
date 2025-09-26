package retry

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewExponentialBackoff_1(t *testing.T) {
	tests := []struct {
		name string
		want *ExponentialBackoff
	}{
		{
			name: "Test case 1",
			want: &ExponentialBackoff{
				InitialInterval: 1 * time.Second,
				MaxBackoffTime:  32 * time.Second,
				Multiplier:      2.0,
				MaxRetryCount:   10,
			},
		},
		{
			name: "Test case 2",
			want: &ExponentialBackoff{
				InitialInterval: 2 * time.Second,
				MaxBackoffTime:  64 * time.Second,
				Multiplier:      3.0,
				MaxRetryCount:   20,
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

			if got := NewExponentialBackoff(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExponentialBackoff() = %v, want %v", got, tt.want)
			}
		})
	}
}
