package web

import (
	"fmt"
	"testing"
)

func TestWithRouterMethods_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerInfo{}
	ctrlInterface := &Controller{}
	ctrlInterface.Init(nil, "", "", nil)
	WithRouterMethods(ctrlInterface, "GET", "POST")(ctrl)

	if len(ctrl.methods) != 2 {
		t.Errorf("Expected 2 methods, got %d", len(ctrl.methods))
	}

	if ctrl.methods["GET"] != "Get" {
		t.Errorf("Expected GET method to be Get, got %s", ctrl.methods["GET"])
	}

	if ctrl.methods["POST"] != "Post" {
		t.Errorf("Expected POST method to be Post, got %s", ctrl.methods["POST"])
	}
}
