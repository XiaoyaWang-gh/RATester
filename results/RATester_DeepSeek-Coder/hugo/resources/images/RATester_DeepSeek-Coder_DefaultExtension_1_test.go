package images

import (
	"fmt"
	"testing"
)

func TestDefaultExtension_1(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want string
	}{
		{
			name: "JPEG",
			f:    JPEG,
			want: ".jpg",
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.DefaultExtension(); got != tt.want {
				t.Errorf("Format.DefaultExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
