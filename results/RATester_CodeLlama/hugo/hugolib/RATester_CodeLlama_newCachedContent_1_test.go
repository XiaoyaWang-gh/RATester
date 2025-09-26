package hugolib

import (
	"fmt"
	"testing"
)

func TestNewCachedContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var m *pageMeta
	var h *HugoSites
	var pi *contentParseInfo
	// When
	c, err := m.newCachedContent(h, pi)
	// Then
	if err != nil {
		t.Errorf("newCachedContent() error = %v", err)
		return
	}
	if c == nil {
		t.Errorf("newCachedContent() c = nil")
		return
	}
}
