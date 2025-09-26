package install

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	srcPath := "testdata/src"
	destPath := "testdata/dest"

	err := CopyDir(srcPath, destPath)
	if err != nil {
		t.Errorf("Expected nil, got %s", err.Error())
	}

	// Test if the source directory exists
	if _, err := os.Stat(srcPath); os.IsNotExist(err) {
		t.Errorf("Source directory does not exist")
	}

	// Test if the destination directory exists
	if _, err := os.Stat(destPath); os.IsNotExist(err) {
		t.Errorf("Destination directory does not exist")
	}

	// Test if the files have been copied correctly
	filepath.Walk(destPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			srcFile, _ := ioutil.ReadFile(path)
			destFile, _ := ioutil.ReadFile(strings.Replace(path, destPath, srcPath, -1))
			if bytes.Compare(srcFile, destFile) != 0 {
				t.Errorf("Files are not copied correctly")
			}
		}
		return nil
	})
}
