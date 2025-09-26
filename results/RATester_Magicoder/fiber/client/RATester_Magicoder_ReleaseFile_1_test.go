package client

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReleaseFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	file := &File{
		reader:    ioutil.NopCloser(strings.NewReader("test")),
		name:      "test.txt",
		fieldName: "file",
		path:      "/tmp/test.txt",
	}

	ReleaseFile(file)

	if file.reader != nil {
		t.Error("Expected file.reader to be nil after ReleaseFile")
	}

	if file.name != "" {
		t.Error("Expected file.name to be empty after ReleaseFile")
	}

	if file.fieldName != "" {
		t.Error("Expected file.fieldName to be empty after ReleaseFile")
	}

	if file.path != "" {
		t.Error("Expected file.path to be empty after ReleaseFile")
	}
}
