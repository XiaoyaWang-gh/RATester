package web

import (
	"fmt"
	"testing"
)

func TestNSCtrlAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f interface{}
	var link LinkNamespace
	link = NSCtrlAny(rootpath, f)
	ns := &Namespace{}
	link(ns)
}
