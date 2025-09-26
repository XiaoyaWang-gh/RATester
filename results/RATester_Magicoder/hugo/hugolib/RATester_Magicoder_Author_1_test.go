package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/gohugoio/hugo/output"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/source"
)

func TestAuthor_1(t *testing.T) {
	type fields struct {
		term                   string
		singular               string
		Staler                 resource.Staler
		params                 *pageMetaParams
		fm                     pageMetaFrontMatter
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
		want   page.Author
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

			p := &pageMeta{
				term:                   tt.fields.term,
				singular:               tt.fields.singular,
				Staler:                 tt.fields.Staler,
				pageMetaParams:         tt.fields.params,
				pageMetaFrontMatter:    tt.fields.fm,
				standaloneOutputFormat: tt.fields.standaloneOutputFormat,
				resourcePath:           tt.fields.resourcePath,
				bundled:                tt.fields.bundled,
				pathInfo:               tt.fields.pathInfo,
				f:                      tt.fields.f,
				content:                tt.fields.content,
				s:                      tt.fields.s,
			}
			if got := p.Author(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pageMeta.Author() = %v, want %v", got, tt.want)
			}
		})
	}
}
