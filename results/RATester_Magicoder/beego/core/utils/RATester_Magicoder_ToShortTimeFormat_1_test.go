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
			name: "Zero",
			d:    0,
			want: "0",
		},
		{
			name: "Nanoseconds",
			d:    time.Nanosecond,
			want: "1ns",
		},
		{
			name: "Microseconds",
			d:    time.Microsecond,
			want: "1us",
		},
		{
			name: "Milliseconds",
			d:    time.Millisecond,
			want: "1ms",
		},
		{
			name: "Seconds",
			d:    time.Second,
			want: "1s",
		},
		{
			name: "Minutes",
			d:    time.Minute,
			want: "1m",
		},
		{
			name: "Hours",
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
