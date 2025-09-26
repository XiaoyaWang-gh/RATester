package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestgetAncestors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
	}

	fs.rootMapToReal.Insert("test/path", []RootMapping{
		{
			From: "test/path",
		},
	})

	roots := fs.getAncestors("test/path")

	if len(roots) != 1 {
		t.Errorf("Expected 1 root, got %d", len(roots))
	}

	if roots[0].key != "test/path" {
		t.Errorf("Expected key 'test/path', got '%s'", roots[0].key)
	}

	if len(roots[0].roots) != 1 {
		t.Errorf("Expected 1 root, got %d", len(roots[0].roots))
	}

	if roots[0].roots[0].From != "test/path" {
		t.Errorf("Expected From 'test/path', got '%s'", roots[0].roots[0].From)
	}
}
