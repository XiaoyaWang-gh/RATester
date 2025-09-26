package images

import (
	"fmt"
	"testing"
)

func TestAbsf32_1(t *testing.T) {
	tests := []struct {
		name string
		x    float32
		want float32
	}{
		{
			name: "positive",
			x:    1.0,
			want: 1.0,
		},
		{
			name: "negative",
			x:    -1.0,
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := absf32(tt.x); got != tt.want {
				t.Errorf("absf32() = %v, want %v", got, tt.want)
			}
		})
	}
}
