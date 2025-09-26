package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetCacheDir_1(t *testing.T) {
	type args struct {
		fs       afero.Fs
		cacheDir string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := GetCacheDir(tt.args.fs, tt.args.cacheDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCacheDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCacheDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
