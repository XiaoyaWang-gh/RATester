package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
)

func TestnewPageMap_1(t *testing.T) {
	type args struct {
		i         int
		s         *Site
		mcache    *dynacache.Cache
		pageTrees *pageTrees
	}
	tests := []struct {
		name string
		args args
		want *pageMap
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

			if got := newPageMap(tt.args.i, tt.args.s, tt.args.mcache, tt.args.pageTrees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPageMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
