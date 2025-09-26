package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
)

func TestMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &dirEntryMeta{
		m: &FileMeta{
			PathInfo: &paths.Path{},
			Name:     "test",
			Filename: "test.txt",
		},
	}

	meta := fi.Meta()

	if meta == nil {
		t.Error("Expected non-nil FileMeta, got nil")
	}

	if meta.Name != "test" {
		t.Errorf("Expected Name to be 'test', got '%s'", meta.Name)
	}

	if meta.Filename != "test.txt" {
		t.Errorf("Expected Filename to be 'test.txt', got '%s'", meta.Filename)
	}
}
