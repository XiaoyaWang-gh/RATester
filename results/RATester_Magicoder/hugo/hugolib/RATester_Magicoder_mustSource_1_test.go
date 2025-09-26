package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/resources/resource"
)

func TestmustSource_1(t *testing.T) {
	type fields struct {
		pm             *pageMap
		StaleInfo      resource.StaleInfo
		shortcodeState *shortcodeHandler
		pi             *contentParseInfo
		enableEmoji    bool
		scopes         *maps.Cache[string, *cachedContentScope]
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
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

			c := &cachedContent{
				pm:             tt.fields.pm,
				StaleInfo:      tt.fields.StaleInfo,
				shortcodeState: tt.fields.shortcodeState,
				pi:             tt.fields.pi,
				enableEmoji:    tt.fields.enableEmoji,
				scopes:         tt.fields.scopes,
			}
			if got := c.mustSource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cachedContent.mustSource() = %v, want %v", got, tt.want)
			}
		})
	}
}
