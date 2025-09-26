package filecache

import (
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"
)

func TestReadOrCreate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	fs := afero.NewMemMapFs()
	c := &Cache{
		Fs: fs,
	}

	id := "test.txt"
	content := "Hello, World!"

	read := func(info ItemInfo, r io.ReadSeeker) error {
		b, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		if string(b) != content {
			return fmt.Errorf("expected %q, got %q", content, string(b))
		}
		return nil
	}

	create := func(info ItemInfo, w io.WriteCloser) error {
		_, err := w.Write([]byte(content))
		if err != nil {
			return err
		}
		return w.Close()
	}

	_, err := c.ReadOrCreate(id, read, create)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.Stat(id)
	if err != nil {
		t.Fatal(err)
	}
}
