package navigation

import (
	"fmt"
	"testing"
)

func TestNewMenuCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := newMenuCache()
	if mc == nil {
		t.Errorf("newMenuCache() = %v, want not nil", mc)
	}
	if mc.m == nil {
		t.Errorf("newMenuCache().m = %v, want not nil", mc.m)
	}
}
