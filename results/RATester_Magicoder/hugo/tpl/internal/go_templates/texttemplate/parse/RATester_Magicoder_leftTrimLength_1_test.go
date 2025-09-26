package parse

import (
	"fmt"
	"testing"
)

func TestleftTrimLength_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want Pos
	}{
		{
			name: "Test case 1",
			s:    "   Hello, World!",
			want: 13,
		},
		{
			name: "Test case 2",
			s:    "Hello, World!   ",
			want: 13,
		},
		{
			name: "Test case 3",
			s:    "   Hello, World!   ",
			want: 13,
		},
		{
			name: "Test case 4",
			s:    "Hello, World!",
			want: 13,
		},
		{
			name: "Test case 5",
			s:    "   ",
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

			if got := leftTrimLength(tt.s); got != tt.want {
				t.Errorf("leftTrimLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
