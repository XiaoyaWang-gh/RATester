package hugofs

import (
	"fmt"
	"sync"
	"testing"

	"github.com/spf13/afero"
)

func TesttrackAndWrapFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &OpenFilesFs{
		Fs:        afero.NewMemMapFs(),
		mu:        sync.Mutex{},
		openFiles: make(map[string]int),
	}

	file, err := fs.Create("test.txt")
	if err != nil {
		t.Fatal(err)
	}

	wrappedFile := fs.trackAndWrapFile(file)

	if _, ok := fs.openFiles["test.txt"]; !ok {
		t.Error("Expected file to be tracked")
	}

	if _, ok := wrappedFile.(*openFilesFsFile); !ok {
		t.Error("Expected wrapped file to be of type *openFilesFsFile")
	}
}
