package filesystems

import (
	"fmt"
	"io"
	"testing"

	"github.com/spf13/afero"
)

func TestPrintFs_2(t *testing.T) {
	type args struct {
		fs     afero.Fs
		path   string
		writer io.Writer
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

			printFs(tt.args.fs, tt.args.path, tt.args.writer)
		})
	}
}
