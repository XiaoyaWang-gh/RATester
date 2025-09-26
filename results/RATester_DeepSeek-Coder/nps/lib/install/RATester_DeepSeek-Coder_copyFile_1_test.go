package install

import (
	"fmt"
	"os"
	"testing"
)

func TestCopyFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	src := "testdata/source.txt"
	dest := "testdata/destination.txt"

	w, err := copyFile(src, dest)
	if err != nil {
		t.Errorf("copyFile() error = %v", err)
		return
	}

	if w <= 0 {
		t.Errorf("copyFile() = %v, want > 0", w)
	}

	_, err = os.Stat(dest)
	if os.IsNotExist(err) {
		t.Errorf("copyFile() destination file does not exist")
	}

	os.Remove(dest)
}
