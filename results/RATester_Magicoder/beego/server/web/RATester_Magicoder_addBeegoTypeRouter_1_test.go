package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestaddBeegoTypeRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	ctrl.addBeegoTypeRouter(reflect.TypeOf(Controller{}), "Get", "GET", "/test")

	if _, ok := ctrl.routers["/test"]; !ok {
		t.Error("Expected router to be added, but it was not")
	}
}
