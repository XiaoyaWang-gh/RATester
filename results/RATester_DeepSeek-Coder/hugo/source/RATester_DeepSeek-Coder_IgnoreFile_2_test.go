package source

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/spf13/afero"
)

func TestIgnoreFile_2(t *testing.T) {
	type fields struct {
		PathSpec      *helpers.PathSpec
		SourceFs      afero.Fs
		shouldInclude func(filename string) bool
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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

			s := &SourceSpec{
				PathSpec:      tt.fields.PathSpec,
				SourceFs:      tt.fields.SourceFs,
				shouldInclude: tt.fields.shouldInclude,
			}
			if got := s.IgnoreFile(tt.args.filename); got != tt.want {
				t.Errorf("SourceSpec.IgnoreFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
