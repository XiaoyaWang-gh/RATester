package cssjs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToArgs_1(t *testing.T) {
	type args struct {
		opts TailwindCSSOptions
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			name: "minify",
			args: args{opts: TailwindCSSOptions{Minify: true}},
			want: []any{"--minify"},
		},
		{
			name: "optimize",
			args: args{opts: TailwindCSSOptions{Optimize: true}},
			want: []any{"--optimize"},
		},
		{
			name: "minify and optimize",
			args: args{opts: TailwindCSSOptions{Minify: true, Optimize: true}},
			want: []any{"--minify", "--optimize"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.opts.toArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
