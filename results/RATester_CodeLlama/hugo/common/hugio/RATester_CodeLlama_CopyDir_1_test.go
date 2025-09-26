package hugio

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCopyDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	from := "/from"
	to := "/to"
	shouldCopy := func(filename string) bool {
		return true
	}
	err := CopyDir(fs, from, to, shouldCopy)
	if err != nil {
		t.Errorf("CopyDir() error = %v", err)
		return
	}
	if _, err := fs.Stat(to); err != nil {
		t.Errorf("CopyDir() error = %v", err)
		return
	}
}
