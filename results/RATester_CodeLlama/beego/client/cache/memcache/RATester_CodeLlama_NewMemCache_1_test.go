package memcache

import (
	"fmt"
	"testing"
)

func TestNewMemCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := NewMemCache()
	if c == nil {
		t.Error("NewMemCache() = nil")
	}
}
