package controllers

import (
	"fmt"
	"testing"
)

func TestSocks5_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	controller := &IndexController{}
	controller.Socks5()

	// Check if the SetInfo and SetType methods were called
	if controller.Data["info"] != "socks5" {
		t.Errorf("Expected SetInfo to set 'socks5', got %v", controller.Data["info"])
	}
	if controller.Data["type"] != "socks5" {
		t.Errorf("Expected SetType to set 'socks5', got %v", controller.Data["type"])
	}

	// Check if the display method was called with the correct argument
	if controller.TplName != "index/list" {
		t.Errorf("Expected display to be called with 'index/list', got %v", controller.TplName)
	}
}
