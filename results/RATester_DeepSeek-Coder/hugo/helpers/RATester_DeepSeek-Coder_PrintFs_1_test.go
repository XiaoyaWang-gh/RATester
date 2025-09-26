package helpers

import (
	"fmt"
	"io"
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
		// TODO: Add test cases.
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
