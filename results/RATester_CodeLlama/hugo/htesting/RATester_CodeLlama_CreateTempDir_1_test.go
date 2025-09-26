package htesting

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCreateTempDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	tempDir, cleanup, err := CreateTempDir(fs, "prefix")
	if err != nil {
		t.Fatal(err)
	}
	defer cleanup()
	if tempDir == "" {
		t.Errorf("tempDir is empty")
	}
	if _, err := fs.Stat(tempDir); err != nil {
		t.Errorf("tempDir does not exist")
	}
}
