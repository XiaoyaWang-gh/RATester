package deps

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/cache/dynacache"
	"github.com/gohugoio/hugo/common/types"
	"github.com/gohugoio/hugo/internal/warpc"
)

func TestClose_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Deps{}
	d.MemCache = &dynacache.Cache{}
	d.WasmDispatchers = &warpc.Dispatchers{}
	d.BuildClosers = &types.Closers{}
	err := d.Close()
	if err != nil {
		t.Errorf("Close() = %v", err)
	}
}
