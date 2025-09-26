package filecache

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

func TestGetOrCreate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "foo"
	create := func() (io.ReadCloser, error) {
		return ioutil.NopCloser(strings.NewReader("bar")), nil
	}

	info, r, err := c.GetOrCreate(id, create)
	if err != nil {
		t.Fatal(err)
	}

	if info.Name != id {
		t.Errorf("expected %q, got %q", id, info.Name)
	}

	if r == nil {
		t.Fatal("expected a reader")
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if string(b) != "bar" {
		t.Errorf("expected %q, got %q", "bar", string(b))
	}
}
