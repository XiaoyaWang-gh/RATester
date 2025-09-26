package output

import (
	"fmt"
	"testing"
)

func TestIsZero_9(t *testing.T) {
	tests := []struct {
		name string
		f    Format
		want bool
	}{
		{
			name: "empty name",
			f:    Format{},
			want: true,
		},
		{
			name: "non-empty name",
			f:    Format{Name: "test"},
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

			if got := tt.f.IsZero(); got != tt.want {
				t.Errorf("Format.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
