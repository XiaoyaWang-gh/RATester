package test

import (
	"fmt"
	"os"
	"testing"
)

func TestviewsHeaderTpl_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a, err := viewsHeaderTpl()
	if err != nil {
		t.Errorf("viewsHeaderTpl() returned an error: %v", err)
	}

	if a == nil {
		t.Error("viewsHeaderTpl() returned nil")
	}

	if a.info == nil {
		t.Error("viewsHeaderTpl() returned nil info")
	}

	if a.info.Name() != "views/header.tpl" {
		t.Errorf("viewsHeaderTpl() returned wrong name: %s", a.info.Name())
	}

	if a.info.Size() != 52 {
		t.Errorf("viewsHeaderTpl() returned wrong size: %d", a.info.Size())
	}

	if a.info.Mode() != os.FileMode(0o664) {
		t.Errorf("viewsHeaderTpl() returned wrong mode: %v", a.info.Mode())
	}

	if a.info.ModTime().Unix() != 1541431067 {
		t.Errorf("viewsHeaderTpl() returned wrong mod time: %v", a.info.ModTime().Unix())
	}
}
