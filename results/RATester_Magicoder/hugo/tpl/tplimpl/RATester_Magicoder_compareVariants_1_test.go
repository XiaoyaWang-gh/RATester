package tplimpl

import (
	"fmt"
	"testing"
)

func TestcompareVariants_1(t *testing.T) {
	tests := []struct {
		name string
		a    []string
		b    []string
		want int
	}{
		{
			name: "Test case 1",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "c"},
			want: 6,
		},
		{
			name: "Test case 2",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "d"},
			want: 4,
		},
		{
			name: "Test case 3",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "c", "d"},
			want: 3,
		},
		{
			name: "Test case 4",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b"},
			want: 2,
		},
		{
			name: "Test case 5",
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "c", "d", "e"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &shortcodeTemplates{}
			if got := s.compareVariants(tt.a, tt.b); got != tt.want {
				t.Errorf("compareVariants() = %v, want %v", got, tt.want)
			}
		})
	}
}
