package babel

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToArgs_2(t *testing.T) {
	type test struct {
		name     string
		opts     Options
		expected []any
	}

	compact := true
	tests := []test{
		{
			name: "all options",
			opts: Options{
				Config:     "config.json",
				Minified:   true,
				NoComments: true,
				Compact:    &compact,
				Verbose:    true,
				NoBabelrc:  true,
				SourceMap:  "external",
			},
			expected: []any{
				"--source-maps",
				"--minified",
				"--no-comments",
				"--compact=true",
				"--verbose",
				"--no-babelrc",
			},
		},
		{
			name: "some options",
			opts: Options{
				Minified:  true,
				Compact:   &compact,
				SourceMap: "inline",
			},
			expected: []any{
				"--source-maps=inline",
				"--minified",
				"--compact=true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.opts.toArgs()
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
