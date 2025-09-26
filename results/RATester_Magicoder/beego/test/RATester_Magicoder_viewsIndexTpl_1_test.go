package test

import (
	"fmt"
	"os"
	"testing"
)

func TestviewsIndexTpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a, err := viewsIndexTpl()
	if err != nil {
		t.Errorf("viewsIndexTpl() returned an error: %v", err)
	}

	if a == nil {
		t.Error("viewsIndexTpl() returned nil")
	}

	if a.info == nil {
		t.Error("viewsIndexTpl() returned nil info")
	}

	if a.info.Name() != "views/index.tpl" {
		t.Errorf("viewsIndexTpl() returned wrong name: %s", a.info.Name())
	}

	if a.info.Size() != 255 {
		t.Errorf("viewsIndexTpl() returned wrong size: %d", a.info.Size())
	}

	if a.info.Mode() != os.FileMode(0o664) {
		t.Errorf("viewsIndexTpl() returned wrong mode: %v", a.info.Mode())
	}

	if a.info.ModTime().Unix() != 1541434906 {
		t.Errorf("viewsIndexTpl() returned wrong mod time: %v", a.info.ModTime().Unix())
	}
}
