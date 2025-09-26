package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/spf13/afero"
)

func TestrealDirOpener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		id: "test",
		Fs: afero.NewMemMapFs(),
	}

	meta := &FileMeta{
		PathInfo: &paths.Path{},
		Name:     "test",
		Filename: "test",
	}

	opener := fs.realDirOpener("test", meta)

	file, err := opener()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if file == nil {
		t.Error("Expected file to be not nil")
	}
}
