package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
	"github.com/spf13/afero"
)

func TestVirtualDirOpener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		id:            "test",
		Fs:            afero.NewMemMapFs(),
		rootMapToReal: radix.New(),
		realMapToRoot: radix.New(),
	}

	name := "testDir"
	opener := fs.virtualDirOpener(name)

	file, err := opener()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if file == nil {
		t.Errorf("Expected a file, got nil")
	}

	if file.Name() != name {
		t.Errorf("Expected file name to be %s, got %s", name, file.Name())
	}
}
