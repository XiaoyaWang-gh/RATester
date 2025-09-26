package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestcreateBeegoRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}
	ctrl.cfg = &Config{}
	ctrl.cfg.WebConfig.Session.SessionOn = true

	route := ctrl.createBeegoRouter(reflect.TypeOf(Controller{}), "/test")

	if route.pattern != "/test" {
		t.Errorf("Expected pattern to be '/test', but got %s", route.pattern)
	}

	if route.routerType != routerTypeBeego {
		t.Errorf("Expected routerType to be routerTypeBeego, but got %d", route.routerType)
	}

	if route.sessionOn != true {
		t.Errorf("Expected sessionOn to be true, but got %t", route.sessionOn)
	}

	if route.controllerType != reflect.TypeOf(Controller{}) {
		t.Errorf("Expected controllerType to be Controller, but got %s", route.controllerType)
	}
}
