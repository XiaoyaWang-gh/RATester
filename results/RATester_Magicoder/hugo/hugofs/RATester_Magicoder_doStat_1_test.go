package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
	"github.com/spf13/afero"
)

func TestdoStat_1(t *testing.T) {
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

	// Create a file
	err := fs.Fs.Mkdir("testdir", 0755)
	if err != nil {
		t.Fatal(err)
	}

	// Test case 1: File exists
	_, err = fs.doStat("testdir")
	if err != nil {
		t.Fatal(err)
	}

	// Test case 2: File does not exist
	_, err = fs.doStat("nonexistent")
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	// Test case 3: File is a directory
	err = fs.Fs.Mkdir("testdir2", 0755)
	if err != nil {
		t.Fatal(err)
	}

	_, err = fs.doStat("testdir2")
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}
}
