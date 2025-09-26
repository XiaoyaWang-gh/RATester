package web

import (
	"fmt"
	"testing"
)

func TestAddToRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	p.routers = make(map[string]*Tree)
	p.cfg = &Config{}
	p.cfg.RouterCaseSensitive = false
	method := "GET"
	pattern := "/"
	r := &ControllerInfo{}
	p.addToRouter(method, pattern, r)
	if len(p.routers) != 1 {
		t.Errorf("TestaddToRouter failed")
	}
}
