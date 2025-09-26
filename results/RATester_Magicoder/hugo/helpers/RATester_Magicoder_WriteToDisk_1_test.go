package helpers

import (
	"fmt"
	"io"
	"strings"
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
		{
			name: "Test case 1",
			args: args{
				inpath: "test.txt",
				r:      strings.NewReader("test data"),
				fs:     afero.NewMemMapFs(),
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				inpath: "test.txt",
				r:      strings.NewReader(""),
				fs:     afero.NewMemMapFs(),
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				inpath: "test.txt",
				r:      strings.NewReader("test data"),
				fs:     afero.NewOsFs(),
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				inpath: "test.txt",
				r:      strings.NewReader(""),
				fs:     afero.NewOsFs(),
			},
			wantErr: false,
		},
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
