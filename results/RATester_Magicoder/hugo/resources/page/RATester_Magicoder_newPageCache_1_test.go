package page

import (
	"fmt"
	"testing"
)

func TestnewPageCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pc := newPageCache()
	if pc == nil {
		t.Error("newPageCache() returned nil")
	}
	if pc.m == nil {
		t.Error("newPageCache() returned a pageCache with nil map")
	}
}
