package filecache

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gohugoio/hugo/common/hugio"
	"github.com/spf13/afero"
)

func TestGetOrCreate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "test_id"
	create := func() (io.ReadCloser, error) {
		return hugio.ToReadCloser(strings.NewReader("test_data")), nil
	}

	info, r, err := cache.GetOrCreate(id, create)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if info.Name != id {
		t.Errorf("Expected name %q, got %q", id, info.Name)
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(data) != "test_data" {
		t.Errorf("Expected data %q, got %q", "test_data", string(data))
	}
}
