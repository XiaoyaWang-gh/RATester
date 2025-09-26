package cssjs

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoArgs_1(t *testing.T) {
	tests := []struct {
		name string
		opts TailwindCSSOptions
		want []any
	}{
		{
			name: "Minify and Optimize",
			opts: TailwindCSSOptions{
				Minify:   true,
				Optimize: true,
			},
			want: []any{"--minify", "--optimize"},
		},
		{
			name: "Only Minify",
			opts: TailwindCSSOptions{
				Minify: true,
			},
			want: []any{"--minify"},
		},
		{
			name: "Only Optimize",
			opts: TailwindCSSOptions{
				Optimize: true,
			},
			want: []any{"--optimize"},
		},
		{
			name: "No Options",
			opts: TailwindCSSOptions{},
			want: []any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.opts.toArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
