package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestPruneRootDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	if _, err := c.pruneRootDir(true); err != nil {
		t.Fatal(err)
	}

	if _, err := c.pruneRootDir(false); err != nil {
		t.Fatal(err)
	}
}
