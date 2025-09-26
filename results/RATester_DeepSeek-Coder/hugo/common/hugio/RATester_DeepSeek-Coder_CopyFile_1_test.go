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
				from: "from",
				to:   "to",
			},
			wantErr: false,
		},
		{
			name: "TestCopyFile_Fail_From_Not_Exist",
			args: args{
				fs:   afero.NewMemMapFs(),
				from: "not_exist",
				to:   "to",
			},
			wantErr: true,
		},
		{
			name: "TestCopyFile_Fail_To_Not_Exist",
			args: args{
				fs:   afero.NewMemMapFs(),
				from: "from",
				to:   "not_exist",
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

			err := CopyFile(tt.args.fs, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("CopyFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
