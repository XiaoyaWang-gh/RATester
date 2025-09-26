package controllers

import (
	"fmt"
	"testing"
)

func TestPrepare_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &BaseController{}
	s.Prepare()
	if s.controllerName != "" {
		t.Errorf("controllerName is %s, want %s", s.controllerName, "")
	}
	if s.actionName != "" {
		t.Errorf("actionName is %s, want %s", s.actionName, "")
	}
}
