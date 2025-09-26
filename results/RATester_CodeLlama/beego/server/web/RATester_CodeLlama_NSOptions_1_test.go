package web

import (
	"fmt"
	"testing"
)

func TestNSOptions_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var rootpath string
	var f HandleFunc
	var link LinkNamespace
	link = NSOptions(rootpath, f)
	ns := &Namespace{}
	link(ns)
}
