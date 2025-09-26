package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestFindRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{
		routers: make(map[string]*Tree),
	}
	context := &beecontext.Context{}
	routerInfo, isFind := p.FindRouter(context)
	if routerInfo != nil {
		t.Errorf("TestFindRouter failed")
	}
	if isFind {
		t.Errorf("TestFindRouter failed")
	}
}
