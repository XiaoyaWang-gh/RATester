package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestImageCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	testCaches := Caches{
		CacheKeyImages: &Cache{
			Fs: afero.NewMemMapFs(),
		},
	}

	expected := testCaches[CacheKeyImages]
	actual := testCaches.ImageCache()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
