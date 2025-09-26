package helpers

import (
	"fmt"
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
		wantErr bool
	}{
		{
			name: "TestOpenFilesForWriting_Success",
			args: args{
				fs:        afero.NewMemMapFs(),
				filenames: []string{"file1.txt", "file2.txt"},
			},
			wantErr: false,
		},
		{
			name: "TestOpenFilesForWriting_Failure",
			args: args{
				fs:        afero.NewMemMapFs(),
				filenames: []string{"file1.txt", "/file2.txt"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := OpenFilesForWriting(tt.args.fs, tt.args.filenames...)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFilesForWriting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
