package wait

import (
	"fmt"
	"testing"
	"time"
)

func TestJitter_1(t *testing.T) {
	tests := []struct {
		name      string
		duration  time.Duration
		maxFactor float64
		want      time.Duration
	}{
		{
			name:      "Test case 1",
			duration:  1 * time.Second,
			maxFactor: 0.5,
			want:      1500 * time.Millisecond,
		},
		{
			name:      "Test case 2",
			duration:  1 * time.Second,
			maxFactor: 0.0,
			want:      1 * time.Second,
		},
		{
			name:      "Test case 3",
			duration:  1 * time.Second,
			maxFactor: 1.0,
			want:      1500 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Jitter(tt.duration, tt.maxFactor); got != tt.want {
				t.Errorf("Jitter() = %v, want %v", got, tt.want)
			}
		})
	}
}
