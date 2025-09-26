package pagesfromdata

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestResolveDeletedPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{}
	b.sourceInfosPrevious = maps.NewCache[string, *sourceInfo]()
	b.sourceInfosCurrent = maps.NewCache[string, *sourceInfo]()
	b.resolveDeletedPaths()
	if len(b.DeletedPaths) != 0 {
		t.Errorf("Expected no deleted paths, got %v", b.DeletedPaths)
	}

	b.sourceInfosPrevious.Set("a", &sourceInfo{hash: 1})
	b.sourceInfosCurrent.Set("a", &sourceInfo{hash: 1})
	b.resolveDeletedPaths()
	if len(b.DeletedPaths) != 0 {
		t.Errorf("Expected no deleted paths, got %v", b.DeletedPaths)
	}

	b.sourceInfosPrevious.Set("a", &sourceInfo{hash: 1})
	b.sourceInfosCurrent.Set("a", &sourceInfo{hash: 2})
	b.resolveDeletedPaths()
	if len(b.DeletedPaths) != 1 {
		t.Errorf("Expected one deleted path, got %v", b.DeletedPaths)
	}

	b.sourceInfosPrevious.Set("a", &sourceInfo{hash: 1})
	b.sourceInfosCurrent.Set("a", &sourceInfo{hash: 1})
	b.sourceInfosPrevious.Set("b", &sourceInfo{hash: 1})
	b.sourceInfosCurrent.Set("b", &sourceInfo{hash: 2})
	b.resolveDeletedPaths()
	if len(b.DeletedPaths) != 2 {
		t.Errorf("Expected two deleted paths, got %v", b.DeletedPaths)
	}
}
