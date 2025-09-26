package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/images"
)

func TestgetImageMetaCacheTargetPath_1(t *testing.T) {
	type fields struct {
		Image *images.Image
		root  *imageResource
		meta  *imageMeta
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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
				Image: tt.fields.Image,
				root:  tt.fields.root,
				meta:  tt.fields.meta,
			}
			if got := i.getImageMetaCacheTargetPath(); got != tt.want {
				t.Errorf("getImageMetaCacheTargetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
