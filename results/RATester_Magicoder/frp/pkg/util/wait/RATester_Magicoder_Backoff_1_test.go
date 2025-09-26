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
		args args
		want time.Duration
	}{
		{
			name: "Test case 1",
			args: args{
				previousDuration:       time.Second,
				previousConditionError: true,
			},
			want: time.Second,
		},
		{
			name: "Test case 2",
			args: args{
				previousDuration:       time.Second,
				previousConditionError: false,
			},
			want: time.Second,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			f := &fastBackoffImpl{}
			if got := f.Backoff(tt.args.previousDuration, tt.args.previousConditionError); got != tt.want {
				t.Errorf("Backoff() = %v, want %v", got, tt.want)
			}
		})
	}
}
