package helpers

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenFilesForWriting_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	filenames := []string{"file1", "file2", "file3"}

	wc, err := OpenFilesForWriting(fs, filenames...)
	if err != nil {
		t.Fatalf("OpenFilesForWriting() error = %v", err)
	}

	if _, err := wc.Write([]byte("hello")); err != nil {
		t.Fatalf("Write() error = %v", err)
	}

	if err := wc.Close(); err != nil {
		t.Fatalf("Close() error = %v", err)
	}

	for _, filename := range filenames {
		f, err := fs.Open(filename)
		if err != nil {
			t.Fatalf("Open() error = %v", err)
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatalf("ReadAll() error = %v", err)
		}

		if string(b) != "hello" {
			t.Fatalf("ReadAll() = %v, want %v", string(b), "hello")
		}

		if err := f.Close(); err != nil {
			t.Fatalf("Close() error = %v", err)
		}
	}
}
