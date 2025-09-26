package cssjs

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/identity"
	"github.com/spf13/afero"
)

func TestnewImportResolver_1(t *testing.T) {
	type args struct {
		r                 io.Reader
		inPath            string
		opts              InlineImports
		fs                afero.Fs
		logger            loggers.Logger
		dependencyManager identity.Manager
	}
	tests := []struct {
		name string
		args args
		want *importResolver
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

			if got := newImportResolver(tt.args.r, tt.args.inPath, tt.args.opts, tt.args.fs, tt.args.logger, tt.args.dependencyManager); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newImportResolver() = %v, want %v", got, tt.want)
			}
		})
	}
}
