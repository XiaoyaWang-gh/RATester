package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugofs/glob"
	"github.com/spf13/afero"
)

func TestNewFilenameFilterFs_1(t *testing.T) {
	type args struct {
		fs     afero.Fs
		base   string
		filter *glob.FilenameFilter
	}
	tests := []struct {
		name string
		args args
		want afero.Fs
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

			if got := newFilenameFilterFs(tt.args.fs, tt.args.base, tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFilenameFilterFs() = %v, want %v", got, tt.want)
			}
		})
	}
}
