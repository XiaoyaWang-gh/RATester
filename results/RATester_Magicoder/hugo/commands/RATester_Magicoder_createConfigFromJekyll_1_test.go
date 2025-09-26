package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/metadecoders"
	"github.com/spf13/afero"
)

func TestcreateConfigFromJekyll_1(t *testing.T) {
	type args struct {
		fs           afero.Fs
		inpath       string
		kind         metadecoders.Format
		jekyllConfig map[string]any
	}
	tests := []struct {
		name    string
		i       *importCommand
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

			if err := tt.i.createConfigFromJekyll(tt.args.fs, tt.args.inpath, tt.args.kind, tt.args.jekyllConfig); (err != nil) != tt.wantErr {
				t.Errorf("importCommand.createConfigFromJekyll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
