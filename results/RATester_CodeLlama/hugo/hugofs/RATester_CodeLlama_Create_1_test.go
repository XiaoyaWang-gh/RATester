package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCreate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &stacktracerFs{
		Fs: afero.NewMemMapFs(),
	}
	f, err := fs.Create("foo")
	if err != nil {
		t.Fatal(err)
	}
	if f == nil {
		t.Fatal("expected a file")
	}
	if _, err := f.WriteString("hello"); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := fs.Stat("foo"); err != nil {
		t.Fatal(err)
	}
}
