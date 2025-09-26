package test

import (
	"fmt"
	"os"
	"testing"
)

func TestViewsIndexTpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a, err := viewsIndexTpl()
	if err != nil {
		t.Fatal(err)
	}
	if a.info.Name() != "views/index.tpl" {
		t.Errorf("Expected %s, got %s", "views/index.tpl", a.info.Name())
	}
	if a.info.Size() != 255 {
		t.Errorf("Expected %d, got %d", 255, a.info.Size())
	}
	if a.info.Mode() != os.FileMode(0o664) {
		t.Errorf("Expected %s, got %s", "0o664", a.info.Mode())
	}
	if a.info.ModTime().Unix() != 1541434906 {
		t.Errorf("Expected %d, got %d", 1541434906, a.info.ModTime().Unix())
	}
}
