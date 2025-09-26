package filecache

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func Testinit_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	c := &Cache{
		Fs: fs,
	}

	err := c.init()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, err = fs.Stat("")
	if err != nil {
		t.Errorf("Expected directory to exist, but got %v", err)
	}
}
