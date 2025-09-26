package hugio

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCopyFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	from := "from"
	to := "to"
	err := CopyFile(fs, from, to)
	if err != nil {
		t.Errorf("CopyFile() error = %v", err)
		return
	}
	if _, err := fs.Stat(to); err != nil {
		t.Errorf("CopyFile() error = %v", err)
		return
	}
}
