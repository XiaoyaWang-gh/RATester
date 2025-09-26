package wait

import (
	"fmt"
	"testing"
	"time"
)

func TestJitter_1(t *testing.T) {
	type args struct {
		duration  time.Duration
		maxFactor float64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
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

			if got := Jitter(tt.args.duration, tt.args.maxFactor); got != tt.want {
				t.Errorf("Jitter() = %v, want %v", got, tt.want)
			}
		})
	}
}
