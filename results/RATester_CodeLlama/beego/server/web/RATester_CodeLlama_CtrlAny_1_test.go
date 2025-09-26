package web

import (
	"fmt"
	"testing"
)

func TestCtrlAny_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	rootpath := "rootpath"
	f := func() {}
	CtrlAny(rootpath, f)
}
