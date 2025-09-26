package resources

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
	"github.com/gohugoio/hugo/resources/images/exif"
)

func TestgetExif_1(t *testing.T) {
	type fields struct {
		Image             *images.Image
		root              *imageResource
		metaInit          sync.Once
		metaInitErr       error
		meta              *imageMeta
		dominantColorInit sync.Once
		dominantColors    []images.Color
		baseResource
	}
	tests := []struct {
		name   string
		fields fields
		want   *exif.ExifInfo
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

			i := &imageResource{
				Image:             tt.fields.Image,
				root:              tt.fields.root,
				metaInit:          tt.fields.metaInit,
				metaInitErr:       tt.fields.metaInitErr,
				meta:              tt.fields.meta,
				dominantColorInit: tt.fields.dominantColorInit,
				dominantColors:    tt.fields.dominantColors,
				baseResource:      tt.fields.baseResource,
			}
			if got := i.getExif(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getExif() = %v, want %v", got, tt.want)
			}
		})
	}
}
