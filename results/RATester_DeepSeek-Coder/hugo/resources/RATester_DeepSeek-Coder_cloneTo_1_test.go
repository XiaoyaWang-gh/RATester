package resources

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/resources/internal"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestCloneTo_1(t *testing.T) {
	type fields struct {
		publishInit          *sync.Once
		sd                   ResourceSourceDescriptor
		paths                internal.ResourcePaths
		sourceFilenameIsHash bool
		h                    *resourceHash // A hash of the source content. Is only calculated in caching situations.
		Staler               resource.Staler
		title                string
		name                 string
		params               map[string]any
		spec                 *Spec
	}
	type args struct {
		targetPath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   resource.Resource
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
			if got := l.cloneTo(tt.args.targetPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericResource.cloneTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
