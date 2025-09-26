package parse

import (
	"fmt"
	"testing"
)

func TesthasRightTrimMarker_1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Test case 1",
			s:    " ",
			want: false,
		},
		{
			name: "Test case 2",
			s:    "  ",
			want: true,
		},
		{
			name: "Test case 3",
			s:    "a ",
			want: false,
		},
		{
			name: "Test case 4",
			s:    " a",
			want: false,
		},
		{
			name: "Test case 5",
			s:    " a ",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := hasRightTrimMarker(tt.s); got != tt.want {
				t.Errorf("hasRightTrimMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
