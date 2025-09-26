package test

import (
	"fmt"
	"os"
	"testing"
)

func TestViewsBlocksBlockTpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a, err := viewsBlocksBlockTpl()
	if err != nil {
		t.Fatal(err)
	}
	if a.info.Name() != "views/blocks/block.tpl" {
		t.Errorf("Expected %s, got %s", "views/blocks/block.tpl", a.info.Name())
	}
	if a.info.Size() != 50 {
		t.Errorf("Expected %d, got %d", 50, a.info.Size())
	}
	if a.info.Mode() != os.FileMode(0o664) {
		t.Errorf("Expected %s, got %s", "0664", a.info.Mode())
	}
	if a.info.ModTime().Unix() != 1541431067 {
		t.Errorf("Expected %d, got %d", 1541431067, a.info.ModTime().Unix())
	}
}
