package commands

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestIsStatic_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &staticSyncer{}
	h := &hugolib.HugoSites{}
	filename := "static/file.txt"
	if !s.isStatic(h, filename) {
		t.Errorf("expected %s to be static", filename)
	}
}
