package web

import (
	"fmt"
	"testing"
)

func TestAddRouterForMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var p ControllerRegister
	var route ControllerInfo
	p.addRouterForMethod(&route)
}
