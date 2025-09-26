package images

import (
	"fmt"
	"testing"
)

func TestRequiresDefaultQuality_1(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want bool
	}{
		{
			name: "JPEG",
			f:    JPEG,
			want: true,
		},
		{
			name: "WEBP",
			f:    WEBP,
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

			if got := tt.f.RequiresDefaultQuality(); got != tt.want {
				t.Errorf("Format.RequiresDefaultQuality() = %v, want %v", got, tt.want)
			}
		})
	}
}
