package web

import (
	"fmt"
	"testing"
)

func TestGetControllerAndAction_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		controllerName: "TestController",
		actionName:     "TestAction",
	}

	controller, action := ctrl.GetControllerAndAction()

	if controller != "TestController" {
		t.Errorf("Expected controller name to be 'TestController', got '%s'", controller)
	}

	if action != "TestAction" {
		t.Errorf("Expected action name to be 'TestAction', got '%s'", action)
	}
}
