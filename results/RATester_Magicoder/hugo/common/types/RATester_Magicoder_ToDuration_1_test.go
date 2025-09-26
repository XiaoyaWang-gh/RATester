package types

import (
	"fmt"
	"testing"
	"time"
)

func TestToDuration_1(t *testing.T) {
	tests := []struct {
		name string
		v    any
		want time.Duration
	}{
		{
			name: "Test case 1",
			v:    "1s",
			want: time.Second,
		},
		{
			name: "Test case 2",
			v:    "1m",
			want: time.Minute,
		},
		{
			name: "Test case 3",
			v:    "1h",
			want: time.Hour,
		},
		{
			name: "Test case 4",
			v:    "1ms",
			want: time.Millisecond,
		},
		{
			name: "Test case 5",
			v:    "1us",
			want: time.Microsecond,
		},
		{
			name: "Test case 6",
			v:    "1ns",
			want: time.Nanosecond,
		},
		{
			name: "Test case 7",
			v:    "1h1m1s",
			want: time.Hour + time.Minute + time.Second,
		},
		{
			name: "Test case 8",
			v:    "1h1m1s1ms1us1ns",
			want: time.Hour + time.Minute + time.Second + time.Millisecond + time.Microsecond + time.Nanosecond,
		},
		{
			name: "Test case 9",
			v:    "invalid",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToDuration(tt.v); got != tt.want {
				t.Errorf("ToDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
