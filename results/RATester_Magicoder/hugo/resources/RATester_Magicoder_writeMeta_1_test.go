package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/cache/filecache"
)

func TestwriteMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &ResourceCache{
		fileCache: &filecache.Cache{},
	}

	key := "testKey"
	meta := transformedResourceMetadata{
		Target:     "testTarget",
		MediaTypeV: "testMediaType",
		MetaData:   map[string]any{"testKey": "testValue"},
	}

	_, _, err := cache.writeMeta(key, meta)
	if err != nil {
		t.Errorf("writeMeta() error = %v", err)
		return
	}

	// Add assertions here to check if the file was written correctly
}
