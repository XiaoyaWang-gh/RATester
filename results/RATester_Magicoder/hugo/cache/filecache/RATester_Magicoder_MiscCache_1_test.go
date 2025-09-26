package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestMiscCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testCaches := Caches{
		CacheKeyMisc: &Cache{
			Fs: afero.NewMemMapFs(),
		},
	}

	cache := testCaches.MiscCache()
	if cache == nil {
		t.Error("Expected non-nil cache")
	}
}
