package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/media"
	"github.com/spf13/afero"
)

func TestToTransformedResourceMetadata_1(t *testing.T) {
	type fields struct {
		content        *string
		sourceFilename *string
		sourceFs       afero.Fs
		targetPath     string
		mediaType      media.Type
		data           map[string]any
		startCtx       ResourceTransformationCtx
	}
	tests := []struct {
		name   string
		fields fields
		want   transformedResourceMetadata
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

			u := &transformationUpdate{
				content:        tt.fields.content,
				sourceFilename: tt.fields.sourceFilename,
				sourceFs:       tt.fields.sourceFs,
				targetPath:     tt.fields.targetPath,
				mediaType:      tt.fields.mediaType,
				data:           tt.fields.data,
				startCtx:       tt.fields.startCtx,
			}
			if got := u.toTransformedResourceMetadata(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toTransformedResourceMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
