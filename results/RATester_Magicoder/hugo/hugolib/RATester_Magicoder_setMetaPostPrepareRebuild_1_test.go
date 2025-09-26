package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/source"
)

func TestsetMetaPostPrepareRebuild_1(t *testing.T) {
	type fields struct {
		term     string
		singular string
		Staler   resource.Staler
		pageMetaParams
		pageMetaFrontMatter
		standaloneOutputFormat output.Format
		resourcePath           string
		bundled                bool
		pathInfo               *paths.Path
		f                      *source.File
		content                *cachedContent
		s                      *Site
	}
	tests := []struct {
		name   string
		fields fields
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

			m := &pageMeta{
				term:     tt.fields.term,
				singular: tt.fields.singular,
				Staler:   tt.fields.Staler,
				pageMetaParams: &pageMetaParams{
					setMetaPostCount:          tt.fields.setMetaPostCount,
					setMetaPostCascadeChanged: tt.fields.setMetaPostCascadeChanged,
					pageConfig:                tt.fields.pageConfig,
					datesOriginal:             tt.fields.datesOriginal,
					paramsOriginal:            tt.fields.paramsOriginal,
					cascadeOriginal:           tt.fields.cascadeOriginal,
				},
				pageMetaFrontMatter:    tt.fields.pageMetaFrontMatter,
				standaloneOutputFormat: tt.fields.standaloneOutputFormat,
				resourcePath:           tt.fields.resourcePath,
				bundled:                tt.fields.bundled,
				pathInfo:               tt.fields.pathInfo,
				f:                      tt.fields.f,
				content:                tt.fields.content,
				s:                      tt.fields.s,
			}
			m.setMetaPostPrepareRebuild()
		})
	}
}
