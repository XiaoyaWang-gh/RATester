package web

import (
	"fmt"
	"testing"
)

func TestNSAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f HandleFunc
	var link LinkNamespace
	link = NSAny(rootpath, f)
	ns := &Namespace{}
	link(ns)
}
