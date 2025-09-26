package hugolib

import (
	"fmt"
	"testing"
)

func TestnewPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	m := &pageMeta{}
	ps, path, err := h.newPage(m)
	if err != nil {
		t.Errorf("newPage returned an error: %v", err)
	}
	if ps == nil {
		t.Error("newPage returned nil pageState")
	}
	if path == nil {
		t.Error("newPage returned nil path")
	}
}
