package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsContent_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := SourceFilesystems{
		Content: &SourceFilesystem{
			Fs: afero.NewMemMapFs(),
		},
	}

	if !s.IsContent("foo.md") {
		t.Errorf("Expected true")
	}

	if s.IsContent("foo.txt") {
		t.Errorf("Expected false")
	}
}
