package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestDecorateFileInfo_1(t *testing.T) {
	type args struct {
		fi       FileNameIsDir
		opener   func() (afero.File, error)
		filename string
		inMeta   *FileMeta
	}
	tests := []struct {
		name string
		args args
		want FileMetaInfo
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

			if got := decorateFileInfo(tt.args.fi, tt.args.opener, tt.args.filename, tt.args.inMeta); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decorateFileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
