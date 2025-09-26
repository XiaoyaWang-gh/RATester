package pagesfromdata

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/identity"
)

func TestPrepareNextBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{
		StaleVersion:        0,
		EnableAllLanguages:  true,
		DeletedPaths:        []string{},
		ChangedIdentities:   []identity.Identity{},
		NumPagesAdded:       0,
		NumResourcesAdded:   0,
		sourceInfosCurrent:  maps.NewCache[string, *sourceInfo](),
		sourceInfosPrevious: maps.NewCache[string, *sourceInfo](),
	}

	b.PrepareNextBuild()

	if b.StaleVersion != 0 {
		t.Errorf("Expected StaleVersion to be 0, but got %d", b.StaleVersion)
	}
	if b.EnableAllLanguages != true {
		t.Errorf("Expected EnableAllLanguages to be true, but got %t", b.EnableAllLanguages)
	}
	if len(b.DeletedPaths) != 0 {
		t.Errorf("Expected DeletedPaths to be empty, but got %v", b.DeletedPaths)
	}
	if len(b.ChangedIdentities) != 0 {
		t.Errorf("Expected ChangedIdentities to be empty, but got %v", b.ChangedIdentities)
	}
	if b.NumPagesAdded != 0 {
		t.Errorf("Expected NumPagesAdded to be 0, but got %d", b.NumPagesAdded)
	}
	if b.NumResourcesAdded != 0 {
		t.Errorf("Expected NumResourcesAdded to be 0, but got %d", b.NumResourcesAdded)
	}
	if b.sourceInfosCurrent == nil {
		t.Errorf("Expected sourceInfosCurrent to not be nil, but got nil")
	}
	if b.sourceInfosPrevious == nil {
		t.Errorf("Expected sourceInfosPrevious to not be nil, but got nil")
	}
}
