package media

import (
	"fmt"
	"testing"
)

func TestIsZero_13(t *testing.T) {
	tests := []struct {
		name string
		m    Type
		want bool
	}{
		{
			name: "empty subtype",
			m:    Type{SubType: ""},
			want: true,
		},
		{
			name: "non-empty subtype",
			m:    Type{SubType: "rss"},
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

			if got := tt.m.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
