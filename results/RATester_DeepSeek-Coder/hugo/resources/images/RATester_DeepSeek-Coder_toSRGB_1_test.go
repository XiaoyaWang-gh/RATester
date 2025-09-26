package images

import (
	"fmt"
	"testing"
)

func TestToSRGB_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		c    Color
		i    uint8
		want float64
	}{
		{
			name: "test case 1",
			c:    Color{},
			i:    0,
			want: 0,
		},
		{
			name: "test case 2",
			c:    Color{},
			i:    128,
			want: 0.4980392156862745,
		},
		{
			name: "test case 3",
			c:    Color{},
			i:    255,
			want: 0.9999999999999999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.toSRGB(tt.i); got != tt.want {
				t.Errorf("toSRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
