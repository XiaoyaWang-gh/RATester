package helpers

import (
	"fmt"
	"io"
	"testing"

	"github.com/spf13/afero"
)

func TestSafeWriteToDisk_1(t *testing.T) {
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

			if err := SafeWriteToDisk(tt.args.inpath, tt.args.r, tt.args.fs); (err != nil) != tt.wantErr {
				t.Errorf("SafeWriteToDisk() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
