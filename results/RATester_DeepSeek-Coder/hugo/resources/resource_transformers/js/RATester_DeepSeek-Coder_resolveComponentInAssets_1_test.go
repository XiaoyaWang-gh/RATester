package js

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
	"github.com/spf13/afero"
)

func TestResolveComponentInAssets_1(t *testing.T) {
	type args struct {
		fs      afero.Fs
		impPath string
	}
	tests := []struct {
		name string
		args args
		want *hugofs.FileMeta
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

			if got := resolveComponentInAssets(tt.args.fs, tt.args.impPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resolveComponentInAssets() = %v, want %v", got, tt.want)
			}
		})
	}
}
