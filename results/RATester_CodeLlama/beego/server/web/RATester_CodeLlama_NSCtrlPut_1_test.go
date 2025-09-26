package web

import (
	"fmt"
	"testing"
)

func TestNSCtrlPut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f interface{}
	var link LinkNamespace
	link = NSCtrlPut(rootpath, f)
	if link == nil {
		t.Errorf("TestNSCtrlPut failed")
	}
}
