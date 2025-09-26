package helpers

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestPrintFs_1(t *testing.T) {
	type args struct {
		fs   afero.Fs
		path string
		w    io.Writer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				fs:   afero.NewMemMapFs(),
				path: "/",
				w:    os.Stdout,
			},
		},
		{
			name: "Test case 2",
			args: args{
				fs:   afero.NewOsFs(),
				path: "/",
				w:    os.Stdout,
			},
		},
		{
			name: "Test case 3",
			args: args{
				fs:   afero.NewBasePathFs(afero.NewOsFs(), "/"),
				path: "/",
				w:    os.Stdout,
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

			PrintFs(tt.args.fs, tt.args.path, tt.args.w)
		})
	}
}
