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
			name: "0",
			d:    0,
			want: "0",
		},
		{
			name: "1ns",
			d:    1,
			want: "1ns",
		},
		{
			name: "1us",
			d:    1000,
			want: "1us",
		},
		{
			name: "1ms",
			d:    1000000,
			want: "1ms",
		},
		{
			name: "1s",
			d:    1000000000,
			want: "1s",
		},
		{
			name: "1m",
			d:    60000000000,
			want: "1m",
		},
		{
			name: "1h",
			d:    3600000000000,
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
