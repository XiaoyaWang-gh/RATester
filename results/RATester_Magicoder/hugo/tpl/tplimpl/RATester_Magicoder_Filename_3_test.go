package tplimpl

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestFilename_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := templateInfo{
		name:     "test",
		template: "test template",
		meta: &hugofs.FileMeta{
			Filename: "test.txt",
		},
	}

	filename := info.Filename()
	if filename != "test.txt" {
		t.Errorf("Expected filename to be 'test.txt', but got '%s'", filename)
	}
}
