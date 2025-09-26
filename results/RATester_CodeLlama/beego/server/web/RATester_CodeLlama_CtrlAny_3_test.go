package web

import (
	"fmt"
	"testing"
)

func TestCtrlAny_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	rootpath := "rootpath"
	f := func() {}
	n.CtrlAny(rootpath, f)
}
