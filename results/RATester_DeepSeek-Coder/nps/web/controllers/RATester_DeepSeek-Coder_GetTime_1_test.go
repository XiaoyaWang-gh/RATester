package controllers

import (
	"fmt"
	"testing"
	"time"

	"github.com/astaxie/beego"
)

func TestGetTime_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &AuthController{
		Controller: beego.Controller{},
	}

	ctrl.GetTime()

	result := ctrl.Data["json"].(map[string]interface{})["time"]
	now := time.Now().Unix()

	if result != now {
		t.Errorf("Expected %v, got %v", now, result)
	}
}
