package web

import (
	"fmt"
	"testing"
)

func TestAddRouterMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{}

	ctrl.AddRouterMethod("GET", "/test", func() {})

	if _, ok := ctrl.routers["GET"]; !ok {
		t.Error("Expected router to be added")
	}
}
