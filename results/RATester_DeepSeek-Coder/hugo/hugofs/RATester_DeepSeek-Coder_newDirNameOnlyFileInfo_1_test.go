package hugofs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewDirNameOnlyFileInfo_1(t *testing.T) {
	type args struct {
		name       string
		meta       *FileMeta
		fileOpener func() (afero.File, error)
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

			got := newDirNameOnlyFileInfo(tt.args.name, tt.args.meta, tt.args.fileOpener)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDirNameOnlyFileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
