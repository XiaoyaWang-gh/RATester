package tplimpl

import (
	"fmt"
	"testing"
)

func TestneedsBaseTemplate_1(t *testing.T) {
	tests := []struct {
		name  string
		templ string
		want  bool
	}{
		{
			name:  "Test case 1",
			templ: "{{/* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 2",
			templ: "{{- /* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 3",
			templ: "{{/* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 4",
			templ: "{{- /* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 5",
			templ: "{{/* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 6",
			templ: "{{- /* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 7",
			templ: "{{/* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 8",
			templ: "{{- /* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 9",
			templ: "{{/* comment */}}",
			want:  true,
		},
		{
			name:  "Test case 10",
			templ: "{{- /* comment */}}",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := needsBaseTemplate(tt.templ); got != tt.want {
				t.Errorf("needsBaseTemplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
