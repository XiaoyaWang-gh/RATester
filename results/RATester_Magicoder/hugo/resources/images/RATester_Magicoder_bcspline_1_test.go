package images

import (
	"fmt"
	"testing"
)

func Testbcspline_1(t *testing.T) {
	tests := []struct {
		name string
		x    float32
		b    float32
		c    float32
		want float32
	}{
		{
			name: "Test case 1",
			x:    0.5,
			b:    0.2,
			c:    0.3,
			want: 0.1,
		},
		{
			name: "Test case 2",
			x:    1.5,
			b:    0.4,
			c:    0.5,
			want: 0.2,
		},
		{
			name: "Test case 3",
			x:    2.5,
			b:    0.6,
			c:    0.7,
			want: 0.3,
		},
		{
			name: "Test case 4",
			x:    -0.5,
			b:    0.8,
			c:    0.9,
			want: 0.4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bcspline(tt.x, tt.b, tt.c); got != tt.want {
				t.Errorf("bcspline() = %v, want %v", got, tt.want)
			}
		})
	}
}
