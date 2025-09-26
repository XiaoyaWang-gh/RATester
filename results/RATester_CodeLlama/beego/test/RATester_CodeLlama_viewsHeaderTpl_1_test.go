package test

import (
	"fmt"
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
		t.Fatal(err)
	}
	if a.info.Name() != "views/header.tpl" {
		t.Errorf("Expected %s, got %s", "views/header.tpl", a.info.Name())
	}
	if a.info.Size() != 52 {
		t.Errorf("Expected %d, got %d", 52, a.info.Size())
	}
	if a.info.Mode() != os.FileMode(0o664) {
		t.Errorf("Expected %s, got %s", "0o664", a.info.Mode())
	}
	if a.info.ModTime().Unix() != 1541431067 {
		t.Errorf("Expected %d, got %d", 1541431067, a.info.ModTime().Unix())
	}
}
