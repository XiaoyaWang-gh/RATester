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
		t.Errorf("NewModelCacheHandler() = nil")
	}
}
