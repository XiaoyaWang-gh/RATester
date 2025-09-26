package filesystems

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestprintFs_2(t *testing.T) {
	type args struct {
		fs   afero.Fs
		path string
		w    io.Writer
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				fs:   afero.NewMemMapFs(),
				path: "/test",
				w:    os.Stdout,
			},
			want: "    \"/test\" \"\"\n",
		},
		{
			name: "Test case 2",
			args: args{
				fs:   afero.NewMemMapFs(),
				path: "/test",
				w:    os.Stdout,
			},
			want: "    \"/test\" \"\"\n",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			printFs(tt.args.fs, tt.args.path, tt.args.w)
			// Add assertions to check the output of the function
		})
	}
}
