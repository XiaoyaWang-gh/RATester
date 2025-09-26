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
	name := "test"
	ctrl.SetInfo(name)
	if ctrl.Data["name"] != name {
		t.Errorf("Expected %s, got %s", name, ctrl.Data["name"])
	}
}
