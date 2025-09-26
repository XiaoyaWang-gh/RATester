package controllers

import (
	"fmt"
	"testing"
)

func TestAjaxErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &BaseController{
		controllerName: "TestController",
	}
	ctrl.AjaxErr("test error")

	data := ctrl.Data["json"].(map[string]interface{})
	if data["msg"] != "test error" {
		t.Errorf("Expected 'test error', got '%v'", data["msg"])
	}
	if data["code"] != 0 {
		t.Errorf("Expected 0, got '%v'", data["code"])
	}
}
