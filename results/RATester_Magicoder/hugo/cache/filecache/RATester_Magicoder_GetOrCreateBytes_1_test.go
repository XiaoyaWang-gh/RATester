package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetOrCreateBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "test_id"
	create := func() ([]byte, error) {
		return []byte("test_data"), nil
	}

	info, data, err := cache.GetOrCreateBytes(id, create)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if info.Name != id {
		t.Errorf("Expected name %s, got %s", id, info.Name)
	}

	if string(data) != "test_data" {
		t.Errorf("Expected data %s, got %s", "test_data", string(data))
	}
}
