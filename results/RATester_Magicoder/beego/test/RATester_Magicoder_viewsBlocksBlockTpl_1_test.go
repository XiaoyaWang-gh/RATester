package test

import (
	"fmt"
	"os"
	"testing"
)

func TestviewsBlocksBlockTpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a, err := viewsBlocksBlockTpl()
	if err != nil {
		t.Errorf("viewsBlocksBlockTpl() returned an error: %v", err)
	}

	if a == nil {
		t.Error("viewsBlocksBlockTpl() returned nil")
	}

	if len(a.bytes) == 0 {
		t.Error("viewsBlocksBlockTpl() returned an empty bytes slice")
	}

	if a.info == nil {
		t.Error("viewsBlocksBlockTpl() returned nil info")
	}

	if a.info.Name() != "views/blocks/block.tpl" {
		t.Errorf("viewsBlocksBlockTpl() returned an incorrect name: %s", a.info.Name())
	}

	if a.info.Size() != 50 {
		t.Errorf("viewsBlocksBlockTpl() returned an incorrect size: %d", a.info.Size())
	}

	if a.info.Mode() != os.FileMode(0o664) {
		t.Errorf("viewsBlocksBlockTpl() returned an incorrect mode: %v", a.info.Mode())
	}

	if a.info.ModTime().Unix() != 1541431067 {
		t.Errorf("viewsBlocksBlockTpl() returned an incorrect mod time: %v", a.info.ModTime().Unix())
	}
}
