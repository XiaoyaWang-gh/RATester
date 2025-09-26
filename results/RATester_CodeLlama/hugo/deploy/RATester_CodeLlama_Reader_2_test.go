package deploy

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"
)

func TestReader_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	lf := &localFile{
		NativePath: "testdata/test.txt",
		fs:         afero.NewOsFs(),
	}

	r, err := lf.Reader()
	if err != nil {
		t.Fatalf("Reader() failed: %v", err)
	}

	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll() failed: %v", err)
	}

	if string(b) != "Hello, world!\n" {
		t.Errorf("ReadAll() = %q, want %q", b, "Hello, world!\n")
	}
}
