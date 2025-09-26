package controllers

import (
	"fmt"
	"testing"
)

func TestAjaxOk_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &BaseController{
		controllerName: "TestController",
	}
	ctrl.AjaxOk("test")

	data := ctrl.Data["json"].(map[string]interface{})
	if data["msg"] != "test" {
		t.Errorf("Expected msg to be 'test', got %s", data["msg"])
	}
	if data["code"] != 1 {
		t.Errorf("Expected code to be 1, got %d", data["code"])
	}
}
