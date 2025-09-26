package web

import (
	"fmt"
	"testing"
)

func TestNSDelete_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f HandleFunc
	var ns *Namespace
	var linkNamespace LinkNamespace
	linkNamespace = NSDelete(rootpath, f)
	ns = &Namespace{}
	linkNamespace(ns)
}
