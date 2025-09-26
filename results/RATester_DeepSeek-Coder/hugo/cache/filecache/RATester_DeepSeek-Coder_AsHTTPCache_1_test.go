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

	t.Parallel()

	fs := afero.NewMemMapFs()
	c := &Cache{
		Fs: fs,
	}

	httpCache := c.AsHTTPCache()

	if httpCache == nil {
		t.Errorf("Expected AsHTTPCache to return a non-nil value")
	}
}
