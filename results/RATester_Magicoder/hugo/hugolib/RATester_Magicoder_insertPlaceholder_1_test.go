package hugolib

import (
	"fmt"
	"testing"
)

func TestinsertPlaceholder_1(t *testing.T) {
	tests := []struct {
		name string
		s    shortcode
		want bool
	}{
		{
			name: "Test case 1",
			s: shortcode{
				doMarkup: true,
			},
			want: false,
		},
		{
			name: "Test case 2",
			s: shortcode{
				doMarkup: false,
			},
			want: true,
		},
		{
			name: "Test case 3",
			s: shortcode{
				doMarkup: true,
			},
			want: false,
		},
		{
			name: "Test case 4",
			s: shortcode{
				doMarkup: false,
			},
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

			if got := tt.s.insertPlaceholder(); got != tt.want {
				t.Errorf("insertPlaceholder() = %v, want %v", got, tt.want)
			}
		})
	}
}
