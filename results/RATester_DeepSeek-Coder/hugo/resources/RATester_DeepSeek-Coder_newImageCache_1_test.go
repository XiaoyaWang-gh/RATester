package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
	"github.com/gohugoio/hugo/cache/filecache"
	"github.com/gohugoio/hugo/helpers"
)

func TestNewImageCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fileCache := &filecache.Cache{}
	memCache := &dynacache.Cache{}
	pathSpec := &helpers.PathSpec{}

	ic := newImageCache(fileCache, memCache, pathSpec)

	if ic.fcache != fileCache {
		t.Errorf("Expected file cache to be %v, got %v", fileCache, ic.fcache)
	}

	if ic.mcache == nil {
		t.Errorf("Expected memory cache to be not nil")
	}

	if ic.pathSpec != pathSpec {
		t.Errorf("Expected path spec to be %v, got %v", pathSpec, ic.pathSpec)
	}
}
