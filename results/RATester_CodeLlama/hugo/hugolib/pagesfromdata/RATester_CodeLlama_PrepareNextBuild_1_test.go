package pagesfromdata

import (
	"fmt"
	"testing"
)

func TestPrepareNextBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{}
	b.PrepareNextBuild()
	if b.StaleVersion != 0 {
		t.Errorf("StaleVersion should be 0, got %d", b.StaleVersion)
	}
	if b.DeletedPaths != nil {
		t.Errorf("DeletedPaths should be nil, got %v", b.DeletedPaths)
	}
	if b.ChangedIdentities != nil {
		t.Errorf("ChangedIdentities should be nil, got %v", b.ChangedIdentities)
	}
	if b.NumPagesAdded != 0 {
		t.Errorf("NumPagesAdded should be 0, got %d", b.NumPagesAdded)
	}
	if b.NumResourcesAdded != 0 {
		t.Errorf("NumResourcesAdded should be 0, got %d", b.NumResourcesAdded)
	}
	if b.sourceInfosCurrent != nil {
		t.Errorf("sourceInfosCurrent should be nil, got %v", b.sourceInfosCurrent)
	}
	if b.sourceInfosPrevious != nil {
		t.Errorf("sourceInfosPrevious should be nil, got %v", b.sourceInfosPrevious)
	}
}
