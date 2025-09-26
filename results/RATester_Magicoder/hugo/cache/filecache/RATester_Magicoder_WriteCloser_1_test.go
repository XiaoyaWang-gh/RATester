package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestWriteCloser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs: fs,
	}

	id := "test.txt"
	info, wc, err := cache.WriteCloser(id)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if info.Name != id {
		t.Errorf("Expected name to be %s, but got %s", id, info.Name)
	}

	if wc == nil {
		t.Error("Expected a WriteCloser, but got nil")
	}

	_, err = wc.Write([]byte("test data"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	err = wc.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, err = fs.Stat(id)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
