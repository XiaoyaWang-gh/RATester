package web

import (
	"fmt"
	"testing"
)

func TestNSCtrlDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f interface{}
	var link LinkNamespace
	link = NSCtrlDelete(rootpath, f)
	ns := &Namespace{}
	link(ns)
}
