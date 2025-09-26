package filesystems

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
	"github.com/spf13/afero"
)

func Testmounts_3(t *testing.T) {
	type fields struct {
		Name          string
		Fs            afero.Fs
		SourceFs      afero.Fs
		PublishFolder string
	}
	tests := []struct {
		name   string
		fields fields
		want   []hugofs.FileMetaInfo
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

			d := &SourceFilesystem{
				Name:          tt.fields.Name,
				Fs:            tt.fields.Fs,
				SourceFs:      tt.fields.SourceFs,
				PublishFolder: tt.fields.PublishFolder,
			}
			if got := d.mounts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SourceFilesystem.mounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
