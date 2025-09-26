package web

import (
	"fmt"
	"testing"
)

func TestNSHead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f HandleFunc
	var link LinkNamespace
	link = NSHead(rootpath, f)
	ns := &Namespace{}
	link(ns)
}
