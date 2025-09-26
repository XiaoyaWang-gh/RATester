package web

import (
	"fmt"
	"testing"
)

func TestOpen_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := FileSystem{}
	name := "test.txt"
	f, err := d.Open(name)
	if err != nil {
		t.Errorf("Open() error = %v", err)
		return
	}
	defer f.Close()
	if _, err := f.Stat(); err != nil {
		t.Errorf("Open() error = %v", err)
		return
	}
	if _, err := f.Readdir(1); err != nil {
		t.Errorf("Open() error = %v", err)
		return
	}
}
