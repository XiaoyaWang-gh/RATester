package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestvirtualDirOpener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	name := "testDir"
	opener := fs.virtualDirOpener(name)

	file, err := opener()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if file == nil {
		t.Error("Expected a file, but got nil")
	}

	if file.Name() != name {
		t.Errorf("Expected file name to be %s, but got %s", name, file.Name())
	}
}
