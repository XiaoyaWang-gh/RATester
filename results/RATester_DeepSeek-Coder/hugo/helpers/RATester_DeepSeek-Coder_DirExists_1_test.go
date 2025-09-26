package helpers

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestDirExists_1(t *testing.T) {
	type args struct {
		path string
		fs   afero.Fs
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "TestDirExists_0",
			args: args{
				path: "/test",
				fs:   afero.NewMemMapFs(),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "TestDirExists_1",
			args: args{
				path: "/test",
				fs:   afero.NewMemMapFs(),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "TestDirExists_2",
			args: args{
				path: "/test",
				fs:   afero.NewOsFs(),
			},
			want:    false,
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

			if tt.args.fs != nil {
				err := tt.args.fs.Mkdir(tt.args.path, 0755)
				if err != nil && !os.IsExist(err) {
					t.Errorf("Error creating directory: %v", err)
				}
			}
			got, err := DirExists(tt.args.path, tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("DirExists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DirExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
