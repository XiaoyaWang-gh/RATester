package resources

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
	"github.com/gohugoio/hugo/cache/filecache"
	"github.com/gohugoio/hugo/helpers"
)

func TestnewImageCache_1(t *testing.T) {
	fileCache := &filecache.Cache{}
	memCache := &dynacache.Cache{}
	pathSpec := &helpers.PathSpec{}

	tests := []struct {
		name string
		want *ImageCache
	}{
		{
			name: "Test Case 1",
			want: newImageCache(fileCache, memCache, pathSpec),
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := newImageCache(fileCache, memCache, pathSpec)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newImageCache() = %v, want %v", got, tt.want)
			}
		})
	}
}
