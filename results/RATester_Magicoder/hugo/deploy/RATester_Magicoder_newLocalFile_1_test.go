package deploy

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestnewLocalFile_1(t *testing.T) {
	type args struct {
		fs         afero.Fs
		nativePath string
		slashpath  string
		m          *deployconfig.Matcher
		mt         media.Types
	}
	tests := []struct {
		name    string
		args    args
		want    *localFile
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

			got, err := newLocalFile(tt.args.fs, tt.args.nativePath, tt.args.slashpath, tt.args.m, tt.args.mt)
			if (err != nil) != tt.wantErr {
				t.Errorf("newLocalFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLocalFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
