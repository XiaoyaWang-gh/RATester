package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestAsHTTPCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	httpCache := cache.AsHTTPCache()

	if httpCache == nil {
		t.Error("Expected httpCache to not be nil")
	}
}
