package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestViewsHeaderTpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a, err := viewsHeaderTpl()
	if err != nil {
		t.Errorf("viewsHeaderTpl() error = %v", err)
		return
	}

	expectedBytes, err := ioutil.ReadFile("views/header.tpl")
	if err != nil {
		t.Errorf("Failed to read expected file: %v", err)
		return
	}

	if !bytes.Equal(a.bytes, expectedBytes) {
		t.Errorf("viewsHeaderTpl() = %v, want %v", a.bytes, expectedBytes)
	}

	expectedInfo, err := os.Stat("views/header.tpl")
	if err != nil {
		t.Errorf("Failed to get file info: %v", err)
		return
	}

	if a.info.Name() != expectedInfo.Name() ||
		a.info.Size() != expectedInfo.Size() ||
		a.info.Mode() != expectedInfo.Mode() ||
		a.info.ModTime() != expectedInfo.ModTime() {
		t.Errorf("viewsHeaderTpl() = %v, want %v", a.info, expectedInfo)
	}
}
