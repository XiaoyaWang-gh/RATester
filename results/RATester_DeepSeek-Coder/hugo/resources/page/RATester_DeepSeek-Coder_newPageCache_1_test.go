package page

import (
	"fmt"
	"testing"
)

func TestNewPageCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pc := newPageCache()
	if pc == nil {
		t.Errorf("newPageCache() = %v, want not nil", pc)
	}
	if pc.m == nil {
		t.Errorf("newPageCache().m = %v, want not nil", pc.m)
	}
}
