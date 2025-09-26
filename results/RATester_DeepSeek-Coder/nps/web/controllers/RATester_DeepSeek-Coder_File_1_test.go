package controllers

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
)

func TestFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &IndexController{
		BaseController: BaseController{
			Controller: beego.Controller{},
		},
	}

	ctrl.File()

	// Check if SetInfo and SetType were called
	if ctrl.Data["info"] != "file server" {
		t.Errorf("Expected SetInfo to set 'file server', got %v", ctrl.Data["info"])
	}
	if ctrl.Data["type"] != "file" {
		t.Errorf("Expected SetType to set 'file', got %v", ctrl.Data["type"])
	}

	// Check if display was called with the correct argument
	if ctrl.TplName != "index/list" {
		t.Errorf("Expected display to be called with 'index/list', got %v", ctrl.TplName)
	}
}
