package tplimpl

import (
	"fmt"
	"testing"
)

func TestCompareVariants_1(t *testing.T) {
	tests := []struct {
		name string
		s    *shortcodeTemplates
		a    []string
		b    []string
		want int
	}{
		{
			name: "Test Case 1",
			s:    &shortcodeTemplates{},
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "c"},
			want: 3,
		},
		{
			name: "Test Case 2",
			s:    &shortcodeTemplates{},
			a:    []string{"a", "b", "c"},
			b:    []string{"a", "b", "d"},
			want: 2,
		},
		{
			name: "Test Case 3",
			s:    &shortcodeTemplates{},
			a:    []string{"a", "b", "c"},
			b:    []string{"d", "e", "f"},
			want: -3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.s.compareVariants(tt.a, tt.b); got != tt.want {
				t.Errorf("shortcodeTemplates.compareVariants() = %v, want %v", got, tt.want)
			}
		})
	}
}
