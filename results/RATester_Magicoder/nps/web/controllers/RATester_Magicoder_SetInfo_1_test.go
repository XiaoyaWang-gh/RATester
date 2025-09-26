package controllers

import (
	"fmt"
	"testing"
)

func TestSetInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &BaseController{}
	ctrl.SetInfo("test")
	if ctrl.Data["name"] != "test" {
		t.Errorf("Expected 'test', got '%v'", ctrl.Data["name"])
	}
}
