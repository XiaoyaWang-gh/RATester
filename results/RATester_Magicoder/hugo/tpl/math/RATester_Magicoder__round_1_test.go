package math

import (
	"fmt"
	"testing"
)

func Test_round_1(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want float64
	}{
		{
			name: "Test 1",
			x:    1.5,
			want: 2,
		},
		{
			name: "Test 2",
			x:    2.5,
			want: 3,
		},
		{
			name: "Test 3",
			x:    3.5,
			want: 4,
		},
		{
			name: "Test 4",
			x:    4.5,
			want: 5,
		},
		{
			name: "Test 5",
			x:    5.5,
			want: 6,
		},
		{
			name: "Test 6",
			x:    6.5,
			want: 7,
		},
		{
			name: "Test 7",
			x:    7.5,
			want: 8,
		},
		{
			name: "Test 8",
			x:    8.5,
			want: 9,
		},
		{
			name: "Test 9",
			x:    9.5,
			want: 10,
		},
		{
			name: "Test 10",
			x:    10.5,
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := _round(tt.x); got != tt.want {
				t.Errorf("_round() = %v, want %v", got, tt.want)
			}
		})
	}
}
