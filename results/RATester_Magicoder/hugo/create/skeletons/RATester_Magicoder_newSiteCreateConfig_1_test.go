package skeletons

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestnewSiteCreateConfig_1(t *testing.T) {
	type args struct {
		fs         afero.Fs
		createpath string
		format     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				fs:         afero.NewMemMapFs(),
				createpath: "/tmp/test",
				format:     "toml",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				fs:         afero.NewMemMapFs(),
				createpath: "/tmp/test",
				format:     "json",
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				fs:         afero.NewMemMapFs(),
				createpath: "/tmp/test",
				format:     "yaml",
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

			if err := newSiteCreateConfig(tt.args.fs, tt.args.createpath, tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("newSiteCreateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
