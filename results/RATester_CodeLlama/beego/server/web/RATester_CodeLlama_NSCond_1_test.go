package web

import (
	"fmt"
	"testing"
)

func TestNSCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var cond namespaceCond
	var link LinkNamespace
	link = NSCond(cond)
	ns := &Namespace{}
	link(ns)
	if ns.Cond(cond) != ns {
		t.Errorf("TestNSCond failed")
	}
}
