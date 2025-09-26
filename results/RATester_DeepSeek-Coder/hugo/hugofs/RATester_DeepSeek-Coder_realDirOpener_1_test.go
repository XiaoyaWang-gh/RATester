package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
	"github.com/spf13/afero"
)

func TestRealDirOpener_1(t *testing.T) {
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
		Filename: "test.txt",
	}

	opener := fs.realDirOpener("test", meta)

	file, err := opener()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if file == nil {
		t.Errorf("Expected a file, got nil")
	}

	_, ok := file.(*rootMappingDir)
	if !ok {
		t.Errorf("Expected a *rootMappingDir, got %T", file)
	}
}
