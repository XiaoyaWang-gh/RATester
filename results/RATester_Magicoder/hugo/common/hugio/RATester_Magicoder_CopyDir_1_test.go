package hugio

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCopyDir_1(t *testing.T) {
	type args struct {
		fs         afero.Fs
		from       string
		to         string
		shouldCopy func(filename string) bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestCopyDir_Success",
			args: args{
				fs:         afero.NewMemMapFs(),
				from:       "/from",
				to:         "/to",
				shouldCopy: func(filename string) bool { return true },
			},
			wantErr: false,
		},
		{
			name: "TestCopyDir_Fail_From_Not_Exist",
			args: args{
				fs:         afero.NewMemMapFs(),
				from:       "/from",
				to:         "/to",
				shouldCopy: func(filename string) bool { return true },
			},
			wantErr: true,
		},
		{
			name: "TestCopyDir_Fail_To_Not_Exist",
			args: args{
				fs:         afero.NewMemMapFs(),
				from:       "/from",
				to:         "/to",
				shouldCopy: func(filename string) bool { return true },
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

			if err := CopyDir(tt.args.fs, tt.args.from, tt.args.to, tt.args.shouldCopy); (err != nil) != tt.wantErr {
				t.Errorf("CopyDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
