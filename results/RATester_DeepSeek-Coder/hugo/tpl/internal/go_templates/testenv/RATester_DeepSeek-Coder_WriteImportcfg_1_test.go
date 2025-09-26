package testenv

import (
	"fmt"
	"testing"
)

func TestWriteImportcfg_1(t *testing.T) {
	type args struct {
		dstPath      string
		packageFiles map[string]string
		pkgs         []string
	}
	tests := []struct {
		name string
		args args
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

			WriteImportcfg(t, tt.args.dstPath, tt.args.packageFiles, tt.args.pkgs...)
		})
	}
}
