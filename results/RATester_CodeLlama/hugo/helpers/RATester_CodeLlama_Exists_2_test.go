package helpers

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestExists_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	path := "test.txt"
	_, err := fs.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	exists, err := Exists(path, fs)
	if err != nil {
		t.Fatal(err)
	}
	if !exists {
		t.Errorf("Expected %s to exist", path)
	}
}
