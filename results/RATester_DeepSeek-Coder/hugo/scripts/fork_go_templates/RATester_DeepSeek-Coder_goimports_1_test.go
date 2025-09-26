package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGoimports_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := "testdata"

	// Create a temporary directory for testing
	tmpDir, err := ioutil.TempDir("", "goimports_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Copy test data to temporary directory
	src, err := os.Open(dir)
	if err != nil {
		t.Fatal(err)
	}
	defer src.Close()

	dst, err := os.OpenFile(tmpDir, os.O_RDONLY|os.O_CREATE|os.O_EXCL|os.O_SYNC, 0777)
	if err != nil {
		t.Fatal(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		t.Fatal(err)
	}

	// Call the function to be tested
	goimports(tmpDir)

	// Check if the files in the temporary directory have been modified
	files, err := ioutil.ReadDir(tmpDir)
	if err != nil {
		t.Fatal(err)
	}

	for _, file := range files {
		original, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			t.Fatal(err)
		}

		modified, err := ioutil.ReadFile(filepath.Join(tmpDir, file.Name()))
		if err != nil {
			t.Fatal(err)
		}

		if bytes.Compare(original, modified) == 0 {
			t.Errorf("File %s was not modified by goimports", file.Name())
		}
	}
}
