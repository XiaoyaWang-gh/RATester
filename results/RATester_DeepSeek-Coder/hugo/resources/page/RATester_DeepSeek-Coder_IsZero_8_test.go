package page

import (
	"fmt"
	"testing"
)

func TestIsZero_8(t *testing.T) {
	tests := []struct {
		name string
		s    Summary
		want bool
	}{
		{
			name: "Empty Summary",
			s:    Summary{},
			want: true,
		},
		{
			name: "Non-empty Summary",
			s:    Summary{Text: "Some text", Type: "auto"},
			want: false,
		},
		{
			name: "Manual Summary",
			s:    Summary{Text: "Some text", Type: "manual"},
			want: false,
		},
		{
			name: "Frontmatter Summary",
			s:    Summary{Text: "Some text", Type: "frontmatter"},
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

			if got := tt.s.IsZero(); got != tt.want {
				t.Errorf("Summary.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
