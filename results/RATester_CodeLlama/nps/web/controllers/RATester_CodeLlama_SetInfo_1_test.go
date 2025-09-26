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

	s := &BaseController{}
	s.SetInfo("name")
	if s.Data["name"] != "name" {
		t.Errorf("SetInfo() = %v, want %v", s.Data["name"], "name")
	}
}
