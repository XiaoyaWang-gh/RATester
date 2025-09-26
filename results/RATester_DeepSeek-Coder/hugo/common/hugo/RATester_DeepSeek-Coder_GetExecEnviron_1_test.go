package hugo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/spf13/afero"
)

func TestGetExecEnviron_1(t *testing.T) {
	type args struct {
		workDir string
		cfg     config.AllProvider
		fs      afero.Fs
	}
	tests := []struct {
		name string
		args args
		want []string
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

			if got := GetExecEnviron(tt.args.workDir, tt.args.cfg, tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetExecEnviron() = %v, want %v", got, tt.want)
			}
		})
	}
}
