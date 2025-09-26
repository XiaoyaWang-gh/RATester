package schema

import (
	"fmt"
	"testing"
)

func TestnewCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := newCache()

	if c == nil {
		t.Error("Expected cache to be not nil")
	}

	if c.m == nil {
		t.Error("Expected cache.m to be not nil")
	}

	if c.regconv == nil {
		t.Error("Expected cache.regconv to be not nil")
	}

	if c.tag != "schema" {
		t.Errorf("Expected cache.tag to be 'schema', but got '%s'", c.tag)
	}
}
