package pagesfromdata

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestcheckHasChangedAndSetSourceInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BuildState{
		sourceInfosCurrent:  maps.NewCache[string, *sourceInfo](),
		sourceInfosPrevious: maps.NewCache[string, *sourceInfo](),
	}

	changedPath := "test.txt"
	v := "test content"

	changed := b.checkHasChangedAndSetSourceInfo(changedPath, v)

	if changed {
		t.Errorf("Expected checkHasChangedAndSetSourceInfo to return false, but got true")
	}

	si, found := b.sourceInfosCurrent.Get(changedPath)
	if !found {
		t.Errorf("Expected sourceInfo to be set in current cache, but not found")
	}

	if si.hash != b.hash(v) {
		t.Errorf("Expected sourceInfo hash to be set correctly, but it's not")
	}

	changedPath2 := "test2.txt"
	v2 := "test content2"

	changed2 := b.checkHasChangedAndSetSourceInfo(changedPath2, v2)

	if !changed2 {
		t.Errorf("Expected checkHasChangedAndSetSourceInfo to return true, but got false")
	}

	si2, found2 := b.sourceInfosCurrent.Get(changedPath2)
	if !found2 {
		t.Errorf("Expected sourceInfo to be set in current cache, but not found")
	}

	if si2.hash != b.hash(v2) {
		t.Errorf("Expected sourceInfo hash to be set correctly, but it's not")
	}
}
