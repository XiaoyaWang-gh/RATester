package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestGofmt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := "testdata"
	gofmt(dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}

		orig, err := ioutil.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		formatted, err := ioutil.ReadFile(filepath.Join(dir, name))
		if err != nil {
			t.Fatalf("Failed to read file: %v", err)
		}

		if bytes.Equal(orig, formatted) {
			t.Errorf("File %s is not formatted", name)
		}
	}
}
