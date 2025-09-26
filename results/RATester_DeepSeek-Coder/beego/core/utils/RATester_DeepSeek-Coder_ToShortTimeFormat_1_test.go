package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestToShortTimeFormat_1(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{
			name: "zero",
			d:    0,
			want: "0",
		},
		{
			name: "nanoseconds",
			d:    time.Nanosecond,
			want: "1ns",
		},
		{
			name: "microseconds",
			d:    time.Microsecond,
			want: "1us",
		},
		{
			name: "milliseconds",
			d:    time.Millisecond,
			want: "1ms",
		},
		{
			name: "seconds",
			d:    time.Second,
			want: "1s",
		},
		{
			name: "minutes",
			d:    time.Minute,
			want: "1m",
		},
		{
			name: "hours",
			d:    time.Hour,
			want: "1h",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToShortTimeFormat(tt.d); got != tt.want {
				t.Errorf("ToShortTimeFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
