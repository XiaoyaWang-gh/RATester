package wait

import (
	"fmt"
	"testing"
	"time"
)

func TestBackoff_1(t *testing.T) {
	type args struct {
		previousDuration       time.Duration
		previousConditionError bool
	}
	tests := []struct {
		name string
		f    *fastBackoffImpl
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

			if got := tt.f.Backoff(tt.args.previousDuration, tt.args.previousConditionError); got != tt.want {
				t.Errorf("fastBackoffImpl.Backoff() = %v, want %v", got, tt.want)
			}
		})
	}
}
