package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestMakeReadableAndRemoveAllModulePkgDir_1(t *testing.T) {
	type args struct {
		fs  afero.Fs
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    int
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

			got, err := MakeReadableAndRemoveAllModulePkgDir(tt.args.fs, tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("MakeReadableAndRemoveAllModulePkgDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MakeReadableAndRemoveAllModulePkgDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
