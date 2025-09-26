package controllers

import (
	"fmt"
	"testing"
)

func TestAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	controller := &IndexController{}
	controller.Init(nil, "IndexController", "All", nil)
	controller.All()

	// Check if the data is set correctly
	if controller.Data["menu"] != "client" {
		t.Errorf("Expected menu to be 'client', got %s", controller.Data["menu"])
	}

	// Check if the client_id is set correctly
	clientId := controller.getEscapeString("client_id")
	if controller.Data["client_id"] != clientId {
		t.Errorf("Expected client_id to be '%s', got %s", clientId, controller.Data["client_id"])
	}

	// Check if the info is set correctly
	expectedInfo := "client id:" + clientId
	if controller.GetSession("info") != expectedInfo {
		t.Errorf("Expected info to be '%s', got %s", expectedInfo, controller.GetSession("info"))
	}

	// Check if the display function is called
	if controller.TplName != "index/list" {
		t.Errorf("Expected TplName to be 'index/list', got %s", controller.TplName)
	}
}
