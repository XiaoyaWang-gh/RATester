package source

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/hugofs/glob"
	"github.com/spf13/afero"
)

func TestNewSourceSpec_1(t *testing.T) {
	type args struct {
		ps              *helpers.PathSpec
		inclusionFilter *glob.FilenameFilter
		fs              afero.Fs
	}
	tests := []struct {
		name string
		args args
		want *SourceSpec
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

			if got := NewSourceSpec(tt.args.ps, tt.args.inclusionFilter, tt.args.fs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSourceSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
