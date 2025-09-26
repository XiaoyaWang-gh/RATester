package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
		t.Errorf("viewsBlocksBlockTpl() error = %v", err)
		return
	}

	expectedBytes, err := ioutil.ReadFile("views/blocks/block.tpl")
	if err != nil {
		t.Errorf("Error reading expected file: %v", err)
		return
	}

	if !bytes.Equal(a.bytes, expectedBytes) {
		t.Errorf("viewsBlocksBlockTpl() = %v, want %v", a.bytes, expectedBytes)
	}

	expectedInfo, err := os.Stat("views/blocks/block.tpl")
	if err != nil {
		t.Errorf("Error getting expected file info: %v", err)
		return
	}

	if a.info.Name() != expectedInfo.Name() ||
		a.info.Size() != expectedInfo.Size() ||
		a.info.Mode() != expectedInfo.Mode() ||
		a.info.ModTime() != expectedInfo.ModTime() {
		t.Errorf("viewsBlocksBlockTpl() = %v, want %v", a.info, expectedInfo)
	}
}
