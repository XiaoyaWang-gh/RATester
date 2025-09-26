package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestloadJekyllConfig_1(t *testing.T) {
	type args struct {
		fs         afero.Fs
		jekyllRoot string
	}
	tests := []struct {
		name string
		c    *importCommand
		args args
		want map[string]any
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

			if got := tt.c.loadJekyllConfig(tt.args.fs, tt.args.jekyllRoot); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadJekyllConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
