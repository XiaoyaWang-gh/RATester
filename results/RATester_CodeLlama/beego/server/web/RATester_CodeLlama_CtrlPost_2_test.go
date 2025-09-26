package web

import (
	"fmt"
	"testing"
)

func TestCtrlPost_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	pattern := "pattern"
	f := func() {}
	p.CtrlPost(pattern, f)
}
