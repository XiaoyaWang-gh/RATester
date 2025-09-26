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
		t.Errorf("Expected a non-nil ModelCache, got nil")
	}

	if mc.cache == nil {
		t.Errorf("Expected a non-nil cache, got nil")
	}

	if mc.cacheByFullName == nil {
		t.Errorf("Expected a non-nil cacheByFullName, got nil")
	}
}
