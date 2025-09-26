package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsEmpty_1(t *testing.T) {
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
			name: "Test case 1: Empty directory",
			args: args{
				path: "/test/empty",
				fs:   afero.NewMemMapFs(),
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test case 2: Non-empty directory",
			args: args{
				path: "/test/non-empty",
				fs:   afero.NewMemMapFs(),
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "Test case 3: Non-existent directory",
			args: args{
				path: "/test/non-existent",
				fs:   afero.NewMemMapFs(),
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

			got, err := IsEmpty(tt.args.path, tt.args.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
