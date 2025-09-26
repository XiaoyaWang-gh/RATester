package tplimpl

import (
	"errors"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TesterrWithFileContext_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := templateInfo{
		name:       "test",
		template:   "test",
		isText:     true,
		isEmbedded: true,
		meta: &hugofs.FileMeta{
			Filename: "test.txt",
		},
	}

	err := info.errWithFileContext("test", errors.New("test error"))
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
