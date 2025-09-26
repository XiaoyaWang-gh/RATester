package htesting

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCreateTempDir_1(t *testing.T) {
	type args struct {
		fs     afero.Fs
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				fs:     afero.NewMemMapFs(),
				prefix: "test",
			},
			want:    "/test",
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				fs:     afero.NewOsFs(),
				prefix: "test",
			},
			want:    "/test",
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				fs:     afero.NewMemMapFs(),
				prefix: "",
			},
			want:    "/",
			wantErr: false,
		},
		{
			name: "Test case 4",
			args: args{
				fs:     afero.NewOsFs(),
				prefix: "",
			},
			want:    "/",
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

			got, cleanup, err := CreateTempDir(tt.args.fs, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTempDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTempDir() = %v, want %v", got, tt.want)
			}
			cleanup()
		})
	}
}
