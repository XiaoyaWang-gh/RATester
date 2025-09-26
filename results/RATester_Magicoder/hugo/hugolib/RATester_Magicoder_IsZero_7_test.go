package hugolib

import (
	"fmt"
	"testing"
)

func TestIsZero_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &cachedContent{
		pi: &contentParseInfo{
			itemsStep2: []any{},
		},
	}

	if !cc.IsZero() {
		t.Error("Expected IsZero to return true")
	}

	cc.pi.itemsStep2 = append(cc.pi.itemsStep2, "test")

	if cc.IsZero() {
		t.Error("Expected IsZero to return false")
	}
}
