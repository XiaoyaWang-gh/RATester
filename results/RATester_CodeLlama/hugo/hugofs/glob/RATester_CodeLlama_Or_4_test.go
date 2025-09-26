package glob

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestOr_4(t *testing.T) {
	tests := []struct {
		name string
		args []glob.Glob
		want glob.Glob
	}{
		{
			name: "empty",
			args: []glob.Glob{},
			want: nil,
		},
		{
			name: "one",
			args: []glob.Glob{glob.MustCompile("a")},
			want: glob.MustCompile("a"),
		},
		{
			name: "two",
			args: []glob.Glob{glob.MustCompile("a"), glob.MustCompile("b")},
			want: globSlice{globs: []glob.Glob{glob.MustCompile("a"), glob.MustCompile("b")}},
		},
		{
			name: "three",
			args: []glob.Glob{glob.MustCompile("a"), glob.MustCompile("b"), glob.MustCompile("c")},
			want: globSlice{globs: []glob.Glob{glob.MustCompile("a"), glob.MustCompile("b"), glob.MustCompile("c")}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Or(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
