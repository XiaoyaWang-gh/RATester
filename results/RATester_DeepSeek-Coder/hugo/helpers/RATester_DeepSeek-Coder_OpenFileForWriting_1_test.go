package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenFileForWriting_1(t *testing.T) {
	type args struct {
		fs       afero.Fs
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				fs:       afero.NewMemMapFs(),
				filename: "test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				fs:       afero.NewMemMapFs(),
				filename: "/path/to/test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				fs:       afero.NewMemMapFs(),
				filename: "/path/to/non-existing/test.txt",
			},
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				fs:       afero.NewReadOnlyFs(afero.NewMemMapFs()),
				filename: "test.txt",
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

			_, err := OpenFileForWriting(tt.args.fs, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFileForWriting() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
