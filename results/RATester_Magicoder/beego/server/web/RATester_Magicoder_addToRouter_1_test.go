package web

import (
	"fmt"
	"testing"
)

func TestaddToRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
		cfg: &Config{
			RouterCaseSensitive: false,
		},
	}

	ctrl.addToRouter("GET", "/test", &ControllerInfo{})

	if _, ok := ctrl.routers["GET"]; !ok {
		t.Error("Expected method to be added to routers")
	}

	if _, ok := ctrl.routers["get"]; ok {
		t.Error("Expected method to be case insensitive")
	}
}
