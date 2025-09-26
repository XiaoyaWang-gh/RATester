package helpers

import (
	"fmt"
	"io"
	"testing"

	"github.com/spf13/afero"
)

func TestWriteToDisk_1(t *testing.T) {
	type args struct {
		inpath string
		r      io.Reader
		fs     afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			if err := WriteToDisk(tt.args.inpath, tt.args.r, tt.args.fs); (err != nil) != tt.wantErr {
				t.Errorf("WriteToDisk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
