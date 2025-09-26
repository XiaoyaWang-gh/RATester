package create

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
	"github.com/spf13/afero"
)

func TestbuildFile_1(t *testing.T) {
	type fields struct {
		archeTypeFs afero.Fs
		sourceFs    afero.Fs
		ps          *helpers.PathSpec
		h           *hugolib.HugoSites
		cf          hugolib.ContentFactory
		archetypeFi hugofs.FileMetaInfo
		targetPath  string
		kind        string
		isDir       bool
		dirMap      archetypeMap
		force       bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
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

			b := &contentBuilder{
				archeTypeFs: tt.fields.archeTypeFs,
				sourceFs:    tt.fields.sourceFs,
				ps:          tt.fields.ps,
				h:           tt.fields.h,
				cf:          tt.fields.cf,
				archetypeFi: tt.fields.archetypeFi,
				targetPath:  tt.fields.targetPath,
				kind:        tt.fields.kind,
				isDir:       tt.fields.isDir,
				dirMap:      tt.fields.dirMap,
				force:       tt.fields.force,
			}
			got, err := b.buildFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("contentBuilder.buildFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("contentBuilder.buildFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
