package hugolib

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/gohugoio/hugo/identity"
	"github.com/gohugoio/hugo/output/layouts"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestgetLayoutDescriptor_1(t *testing.T) {
	type fields struct {
		pid                               uint64
		pageOutputs                       []*pageOutput
		pageOutputTemplateVariationsState *atomic.Uint32
		pageOutputIdx                     int
		pageOutput                        *pageOutput
		pageCommon                        *pageCommon
		Staler                            resource.Staler
		dependencyManager                 identity.Manager
		resourcesPublishInit              *sync.Once
	}
	tests := []struct {
		name   string
		fields fields
		want   layouts.LayoutDescriptor
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

			p := &pageState{
				pid:                               tt.fields.pid,
				pageOutputs:                       tt.fields.pageOutputs,
				pageOutputTemplateVariationsState: tt.fields.pageOutputTemplateVariationsState,
				pageOutputIdx:                     tt.fields.pageOutputIdx,
				pageOutput:                        tt.fields.pageOutput,
				pageCommon:                        tt.fields.pageCommon,
				Staler:                            tt.fields.Staler,
				dependencyManager:                 tt.fields.dependencyManager,
				resourcesPublishInit:              tt.fields.resourcesPublishInit,
			}
			if got := p.getLayoutDescriptor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLayoutDescriptor() = %v, want %v", got, tt.want)
			}
		})
	}
}
