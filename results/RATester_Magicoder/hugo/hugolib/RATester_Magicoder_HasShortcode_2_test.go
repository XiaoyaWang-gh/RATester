package hugolib

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/gohugoio/hugo/identity"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestHasShortcode_2(t *testing.T) {
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
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
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
			if got := p.HasShortcode(tt.args.name); got != tt.want {
				t.Errorf("HasShortcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
