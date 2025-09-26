package navigation

import (
	"fmt"
	"testing"
)

func TestnewMenuCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := newMenuCache()
	if mc == nil {
		t.Error("newMenuCache() should not return nil")
	}
	if mc.m == nil {
		t.Error("newMenuCache() should initialize m")
	}
}
