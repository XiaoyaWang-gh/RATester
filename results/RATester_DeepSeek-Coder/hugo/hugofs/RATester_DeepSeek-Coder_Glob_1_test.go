package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGlob_1(t *testing.T) {
	type args struct {
		fs      afero.Fs
		pattern string
		handle  func(fi FileMetaInfo) (bool, error)
	}
	tests := []struct {
		name    string
		args    args
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

			err := Glob(tt.args.fs, tt.args.pattern, tt.args.handle)
			if (err != nil) != tt.wantErr {
				t.Errorf("Glob() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
