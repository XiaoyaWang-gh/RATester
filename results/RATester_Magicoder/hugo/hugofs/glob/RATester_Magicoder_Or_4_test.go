package glob

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gobwas/glob"
)

func TestOr_4(t *testing.T) {
	type args struct {
		globs []glob.Glob
	}
	tests := []struct {
		name string
		args args
		want glob.Glob
	}{
		{
			name: "Test case 1",
			args: args{
				globs: []glob.Glob{
					glob.MustCompile("*"),
					glob.MustCompile("*"),
				},
			},
			want: globSlice{
				globs: []glob.Glob{
					glob.MustCompile("*"),
					glob.MustCompile("*"),
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				globs: []glob.Glob{
					glob.MustCompile("test*"),
					glob.MustCompile("*test"),
				},
			},
			want: globSlice{
				globs: []glob.Glob{
					glob.MustCompile("test*"),
					glob.MustCompile("*test"),
				},
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

			if got := Or(tt.args.globs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}
