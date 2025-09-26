package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "foo"
	c.WriteCloser(id)

	if got := c.GetString(id); got != "" {
		t.Errorf("GetString() = %v, want %v", got, "")
	}
}
