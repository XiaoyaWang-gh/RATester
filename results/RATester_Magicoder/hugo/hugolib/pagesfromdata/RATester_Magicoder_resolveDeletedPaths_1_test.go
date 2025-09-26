package pagesfromdata

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestresolveDeletedPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{
		sourceInfosPrevious: maps.NewCache[string, *sourceInfo](),
		sourceInfosCurrent:  maps.NewCache[string, *sourceInfo](),
	}

	b.sourceInfosPrevious.Set("path1", &sourceInfo{hash: 1})
	b.sourceInfosPrevious.Set("path2", &sourceInfo{hash: 2})
	b.sourceInfosCurrent.Set("path1", &sourceInfo{hash: 1})
	b.sourceInfosCurrent.Set("path3", &sourceInfo{hash: 3})

	b.resolveDeletedPaths()

	if len(b.DeletedPaths) != 1 || b.DeletedPaths[0] != "path2" {
		t.Errorf("Expected deleted paths to be ['path2'], but got %v", b.DeletedPaths)
	}
}
