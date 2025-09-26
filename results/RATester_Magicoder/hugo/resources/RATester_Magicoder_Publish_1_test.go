package resources

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestPublish_1(t *testing.T) {
	type fields struct {
		publishInit          *sync.Once
		sd                   ResourceSourceDescriptor
		paths                internal.ResourcePaths
		sourceFilenameIsHash bool
		h                    *resourceHash
		Staler               resource.Staler
		title                string
		name                 string
		params               map[string]any
		spec                 *Spec
	}
	tests := []struct {
		name    string
		fields  fields
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

			l := &genericResource{
				publishInit:          tt.fields.publishInit,
				sd:                   tt.fields.sd,
				paths:                tt.fields.paths,
				sourceFilenameIsHash: tt.fields.sourceFilenameIsHash,
				h:                    tt.fields.h,
				Staler:               tt.fields.Staler,
				title:                tt.fields.title,
				name:                 tt.fields.name,
				params:               tt.fields.params,
				spec:                 tt.fields.spec,
			}
			if err := l.Publish(); (err != nil) != tt.wantErr {
				t.Errorf("genericResource.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
