package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestDelete_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}
	httpCache := &httpCache{
		c: cache,
	}

	key := "testKey"
	cache.Fs.Create(key)

	httpCache.Delete(key)

	_, err := cache.Fs.Stat(key)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
