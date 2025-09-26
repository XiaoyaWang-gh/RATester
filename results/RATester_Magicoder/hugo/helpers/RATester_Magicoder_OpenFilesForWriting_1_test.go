package helpers

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenFilesForWriting_1(t *testing.T) {
	type args struct {
		fs        afero.Fs
		filenames []string
	}
	tests := []struct {
		name    string
		args    args
		want    io.WriteCloser
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

			got, err := OpenFilesForWriting(tt.args.fs, tt.args.filenames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFilesForWriting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenFilesForWriting() = %v, want %v", got, tt.want)
			}
		})
	}
}
