package hugofs

import (
	"fmt"
	"os"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &stacktracerFs{
		Fs: afero.NewMemMapFs(),
	}
	f, err := fs.OpenFile("test", os.O_RDONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	if f == nil {
		t.Fatal("expected a file")
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}
