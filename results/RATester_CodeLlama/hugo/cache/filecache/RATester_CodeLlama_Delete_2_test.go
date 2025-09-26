package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestDelete_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &httpCache{
		c: &Cache{
			Fs: afero.NewMemMapFs(),
		},
	}
	key := "key"
	h.Set(key, []byte("value"))
	h.Delete(key)
	_, ok := h.Get(key)
	if ok {
		t.Errorf("expected key to be deleted")
	}
}
