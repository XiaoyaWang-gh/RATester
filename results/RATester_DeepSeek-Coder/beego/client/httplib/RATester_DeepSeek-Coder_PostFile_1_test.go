package httplib

import (
	"fmt"
	"testing"
)

func TestPostFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{
		url:   "http://example.com",
		files: make(map[string]string),
	}

	formname := "file"
	filename := "test.txt"

	b.PostFile(formname, filename)

	if _, ok := b.files[formname]; !ok {
		t.Errorf("Expected file %s to be added to the files map", formname)
	}

	if b.files[formname] != filename {
		t.Errorf("Expected filename %s for formname %s, got %s", filename, formname, b.files[formname])
	}
}
