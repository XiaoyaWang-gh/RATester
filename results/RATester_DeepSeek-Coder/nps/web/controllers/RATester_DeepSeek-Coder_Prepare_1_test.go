package controllers

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
)

func TestPrepare_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &BaseController{
		controllerName: "test",
		actionName:     "test",
	}
	ctrl.Prepare()

	// Check if web_base_url is set correctly
	if ctrl.Data["web_base_url"] != beego.AppConfig.String("web_base_url") {
		t.Errorf("Expected web_base_url to be %s, got %s", beego.AppConfig.String("web_base_url"), ctrl.Data["web_base_url"])
	}

	// Check if controllerName and actionName are set correctly
	if ctrl.controllerName != "test" {
		t.Errorf("Expected controllerName to be 'test', got %s", ctrl.controllerName)
	}
	if ctrl.actionName != "test" {
		t.Errorf("Expected actionName to be 'test', got %s", ctrl.actionName)
	}

	// Check if isAdmin is set correctly
	if ctrl.Data["isAdmin"] != true {
		t.Errorf("Expected isAdmin to be true, got %v", ctrl.Data["isAdmin"])
	}

	// Check if other config values are set correctly
	configValues := []string{
		"https_just_proxy",
		"allow_user_login",
		"allow_flow_limit",
		"allow_rate_limit",
		"allow_connection_num_limit",
		"allow_multi_ip",
		"system_info_display",
		"allow_tunnel_num_limit",
		"allow_local_proxy",
		"allow_user_change_username",
	}
	for _, key := range configValues {
		value, err := beego.AppConfig.Bool(key)
		if err != nil {
			t.Errorf("Error getting config value for %s: %v", key, err)
		}
		if ctrl.Data[key] != value {
			t.Errorf("Expected %s to be %v, got %v", key, value, ctrl.Data[key])
		}
	}
}
