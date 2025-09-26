package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRestoreAssets_1(t *testing.T) {
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tests := []struct {
		name    string
		dir     string
		wantErr bool
	}{
		{
			name:    "valid directory",
			dir:     "valid_directory",
			wantErr: false,
		},
		{
			name:    "invalid directory",
			dir:     "invalid_directory",
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

			err := RestoreAssets(filepath.Join(dir, tt.dir), tt.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestoreAssets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
