package hugio

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCopyFile_1(t *testing.T) {
	type args struct {
		fs   afero.Fs
		from string
		to   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestCopyFile_Success",
			args: args{
				fs:   afero.NewMemMapFs(),
				from: "test.txt",
				to:   "test_copy.txt",
			},
			wantErr: false,
		},
		{
			name: "TestCopyFile_Fail_FromFileNotFound",
			args: args{
				fs:   afero.NewMemMapFs(),
				from: "test_not_found.txt",
				to:   "test_copy.txt",
			},
			wantErr: true,
		},
		{
			name: "TestCopyFile_Fail_ToFileExists",
			args: args{
				fs:   afero.NewMemMapFs(),
				from: "test.txt",
				to:   "test_copy.txt",
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

			if err := CopyFile(tt.args.fs, tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
