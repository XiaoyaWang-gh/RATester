package tableofcontents

import (
	"fmt"
	"testing"
)

func TestIsZero_2(t *testing.T) {
	tests := []struct {
		name string
		h    Heading
		want bool
	}{
		{
			name: "empty",
			h:    Heading{},
			want: true,
		},
		{
			name: "id",
			h:    Heading{ID: "id"},
			want: false,
		},
		{
			name: "title",
			h:    Heading{Title: "title"},
			want: false,
		},
		{
			name: "both",
			h:    Heading{ID: "id", Title: "title"},
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

			if got := tt.h.IsZero(); got != tt.want {
				t.Errorf("Heading.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
