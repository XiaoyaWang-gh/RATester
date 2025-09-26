package filecache

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestRemoveIfExpired_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "foo"

	// Create a file with a modification time in the past.
	f, err := c.Fs.Create(id)
	if err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	if err := c.Fs.Chtimes(id, time.Now().Add(-time.Hour), time.Now().Add(-time.Hour)); err != nil {
		t.Fatal(err)
	}

	if _, err := c.removeIfExpired(id); err != nil {
		t.Fatal(err)
	}

	if _, err := c.Fs.Stat(id); !os.IsNotExist(err) {
		t.Fatalf("expected file to be removed: %v", err)
	}
}
