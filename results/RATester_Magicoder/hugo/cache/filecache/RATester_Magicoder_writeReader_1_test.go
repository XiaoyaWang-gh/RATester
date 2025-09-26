package filecache

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

func TestwriteReader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	c := &Cache{Fs: fs}

	id := "test.txt"
	r := strings.NewReader("test data")

	err := c.writeReader(id, r)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	exists, err := afero.Exists(fs, id)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !exists {
		t.Errorf("Expected file %s to exist", id)
	}

	data, err := afero.ReadFile(fs, id)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if string(data) != "test data" {
		t.Errorf("Expected data to be 'test data', but got %s", string(data))
	}
}
