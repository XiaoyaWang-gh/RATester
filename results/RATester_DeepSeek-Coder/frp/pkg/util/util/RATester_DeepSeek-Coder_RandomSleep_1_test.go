package util

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomSleep_1(t *testing.T) {
	type args struct {
		duration time.Duration
		minRatio float64
		maxRatio float64
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

			if got := RandomSleep(tt.args.duration, tt.args.minRatio, tt.args.maxRatio); got != tt.want {
				t.Errorf("RandomSleep() = %v, want %v", got, tt.want)
			}
		})
	}
}
