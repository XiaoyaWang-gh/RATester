package cssjs

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/spf13/afero"
)

func TestResolve_1(t *testing.T) {
	type fields struct {
		r      io.Reader
		inPath string
		opts   InlineImports
		fs     afero.Fs
		logger loggers.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		want    io.Reader
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

			imp := &importResolver{
				r:      tt.fields.r,
				inPath: tt.fields.inPath,
				opts:   tt.fields.opts,
				fs:     tt.fields.fs,
				logger: tt.fields.logger,
			}
			got, err := imp.resolve()
			if (err != nil) != tt.wantErr {
				t.Errorf("importResolver.resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importResolver.resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
