package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetTempDir_1(t *testing.T) {
	type args struct {
		subPath string
		fs      afero.Fs
	}
	tests := []struct {
		name string
		args args
		want string
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

			if got := GetTempDir(tt.args.subPath, tt.args.fs); got != tt.want {
				t.Errorf("GetTempDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
