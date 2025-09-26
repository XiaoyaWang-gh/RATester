package hreflect

import (
	"fmt"
	"testing"
)

func TestIsTruthful_1(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want bool
	}{
		{
			name: "Test case 1",
			in:   "Hello, World!",
			want: true,
		},
		{
			name: "Test case 2",
			in:   0,
			want: false,
		},
		{
			name: "Test case 3",
			in:   "",
			want: false,
		},
		{
			name: "Test case 4",
			in:   nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsTruthful(tt.in); got != tt.want {
				t.Errorf("IsTruthful() = %v, want %v", got, tt.want)
			}
		})
	}
}
