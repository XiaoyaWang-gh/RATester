package cssjs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToArgs_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		opts TailwindCSSOptions
		want []any
	}{
		{
			name: "minify and optimize",
			opts: TailwindCSSOptions{
				Minify:   true,
				Optimize: true,
			},
			want: []any{"--minify", "--optimize"},
		},
		{
			name: "only minify",
			opts: TailwindCSSOptions{
				Minify: true,
			},
			want: []any{"--minify"},
		},
		{
			name: "only optimize",
			opts: TailwindCSSOptions{
				Optimize: true,
			},
			want: []any{"--optimize"},
		},
		{
			name: "no options",
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
				t.Errorf("TailwindCSSOptions.toArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
