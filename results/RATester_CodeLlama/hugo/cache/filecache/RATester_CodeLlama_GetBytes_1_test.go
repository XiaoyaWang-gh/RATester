package filecache

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestGetBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Cache{
		Fs: afero.NewMemMapFs(),
	}

	id := "foo"
	b := []byte("bar")
	if err := c.writeReader(id, bytes.NewReader(b)); err != nil {
		t.Fatal(err)
	}

	info, got, err := c.GetBytes(id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(info, ItemInfo{Name: id}) {
		t.Errorf("expected %v, got %v", ItemInfo{Name: id}, info)
	}

	if !reflect.DeepEqual(got, b) {
		t.Errorf("expected %v, got %v", b, got)
	}
}
