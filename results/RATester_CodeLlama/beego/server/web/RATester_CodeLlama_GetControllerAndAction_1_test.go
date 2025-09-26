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

	c := &Controller{}
	c.controllerName = "controllerName"
	c.actionName = "actionName"
	controllerName, actionName := c.GetControllerAndAction()
	if controllerName != "controllerName" {
		t.Errorf("controllerName is not equal to controllerName")
	}
	if actionName != "actionName" {
		t.Errorf("actionName is not equal to actionName")
	}
}
