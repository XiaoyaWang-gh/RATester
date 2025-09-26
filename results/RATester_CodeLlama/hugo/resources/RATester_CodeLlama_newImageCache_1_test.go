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

	// TODO: Setup the test case
	fileCache := &filecache.Cache{}
	memCache := &dynacache.Cache{}
	ps := &helpers.PathSpec{}
	// TODO: Call the method under test
	got := newImageCache(fileCache, memCache, ps)
	// TODO: Add test cases
	if got == nil {
		t.Errorf("newImageCache() = nil, want not nil")
	}
}
