package try

import (
	"fmt"
	"testing"
	"time"
)

func TestApplyCIMultiplier_1(t *testing.T) {
	type args struct {
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "Test case 1",
			args: args{timeout: 1 * time.Second},
			want: 1 * time.Second,
		},
		{
			name: "Test case 2",
			args: args{timeout: 2 * time.Second},
			want: 2 * time.Second,
		},
		{
			name: "Test case 3",
			args: args{timeout: 3 * time.Second},
			want: 3 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := applyCIMultiplier(tt.args.timeout); got != tt.want {
				t.Errorf("applyCIMultiplier() = %v, want %v", got, tt.want)
			}
		})
	}
}
