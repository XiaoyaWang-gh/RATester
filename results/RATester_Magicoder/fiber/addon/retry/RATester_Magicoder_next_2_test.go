package retry

import (
	"fmt"
	"testing"
	"time"
)

func Testnext_2(t *testing.T) {
	e := &ExponentialBackoff{
		InitialInterval: 1 * time.Second,
		MaxBackoffTime:  10 * time.Second,
		Multiplier:      2,
		currentInterval: 1 * time.Second,
	}

	tests := []struct {
		name string
		e    *ExponentialBackoff
		want time.Duration
	}{
		{
			name: "Test case 1",
			e:    e,
			want: 2 * time.Second,
		},
		{
			name: "Test case 2",
			e:    e,
			want: 4 * time.Second,
		},
		{
			name: "Test case 3",
			e:    e,
			want: 8 * time.Second,
		},
		{
			name: "Test case 4",
			e:    e,
			want: 10 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.next(); got != tt.want {
				t.Errorf("ExponentialBackoff.next() = %v, want %v", got, tt.want)
			}
		})
	}
}
