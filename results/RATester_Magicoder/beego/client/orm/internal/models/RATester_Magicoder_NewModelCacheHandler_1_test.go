package models

import (
	"fmt"
	"testing"
)

func TestNewModelCacheHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	mc := NewModelCacheHandler()

	if mc == nil {
		t.Error("Expected a non-nil ModelCache, but got nil")
	}

	if mc.cache == nil {
		t.Error("Expected a non-nil cache, but got nil")
	}

	if mc.cacheByFullName == nil {
		t.Error("Expected a non-nil cacheByFullName, but got nil")
	}
}
