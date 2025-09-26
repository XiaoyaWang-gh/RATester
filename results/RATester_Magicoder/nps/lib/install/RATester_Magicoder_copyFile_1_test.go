package install

import (
	"fmt"
	"os"
	"testing"
)

func TestcopyFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	src := "testdata/source.txt"
	dest := "testdata/destination.txt"

	_, err := copyFile(src, dest)
	if err != nil {
		t.Errorf("Test failed, expected nil, got %s", err.Error())
	}

	_, err = os.Stat(dest)
	if os.IsNotExist(err) {
		t.Errorf("Test failed, expected file %s to exist, but it does not", dest)
	}

	srcFile, _ := os.Open(src)
	srcFileInfo, _ := srcFile.Stat()
	srcFileSize := srcFileInfo.Size()

	destFile, _ := os.Open(dest)
	destFileInfo, _ := destFile.Stat()
	destFileSize := destFileInfo.Size()

	if srcFileSize != destFileSize {
		t.Errorf("Test failed, expected file sizes to be equal, got %d and %d", srcFileSize, destFileSize)
	}

	os.Remove(dest)
}
