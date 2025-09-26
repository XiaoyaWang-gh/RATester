package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestGetAncestors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		rootMapToReal: radix.New(),
	}

	fs.rootMapToReal.Insert("/a/b/c", []RootMapping{
		{From: "/a/b/c"},
		{From: "/a/b"},
		{From: "/a"},
	})

	roots := fs.getAncestors("/a/b/c")

	if len(roots) != 3 {
		t.Errorf("Expected 3 roots, got %d", len(roots))
	}

	if roots[0].key != "/a/b/c" {
		t.Errorf("Expected key '/a/b/c', got %s", roots[0].key)
	}

	if len(roots[0].roots) != 1 {
		t.Errorf("Expected 1 root in first key, got %d", len(roots[0].roots))
	}

	if roots[1].key != "/a/b" {
		t.Errorf("Expected key '/a/b', got %s", roots[1].key)
	}

	if len(roots[1].roots) != 1 {
		t.Errorf("Expected 1 root in second key, got %d", len(roots[1].roots))
	}

	if roots[2].key != "/a" {
		t.Errorf("Expected key '/a', got %s", roots[2].key)
	}

	if len(roots[2].roots) != 1 {
		t.Errorf("Expected 1 root in third key, got %d", len(roots[2].roots))
	}
}
