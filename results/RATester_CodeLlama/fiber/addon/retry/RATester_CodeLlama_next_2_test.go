package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestNext_2(t *testing.T) {
	type args struct {
		e *ExponentialBackoff
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "test next",
			args: args{
				e: &ExponentialBackoff{
					InitialInterval: 100 * time.Millisecond,
					MaxBackoffTime:  10 * time.Second,
					Multiplier:      1.5,
					MaxRetryCount:   10,
				},
			},
			want: 100 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.e.next(); got != tt.want {
				t.Errorf("next() = %v, want %v", got, tt.want)
			}
		})
	}
}
