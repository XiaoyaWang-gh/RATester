package controllers

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
)

func TestError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &BaseController{}
	ctrl.error()

	// Check if the layout is set correctly
	if ctrl.Layout != "public/layout.html" {
		t.Errorf("Expected layout to be 'public/layout.html', got '%s'", ctrl.Layout)
	}

	// Check if the template name is set correctly
	if ctrl.TplName != "public/error.html" {
		t.Errorf("Expected template name to be 'public/error.html', got '%s'", ctrl.TplName)
	}

	// Check if the web_base_url is set correctly
	if ctrl.Data["web_base_url"] != beego.AppConfig.String("web_base_url") {
		t.Errorf("Expected web_base_url to be '%s', got '%s'", beego.AppConfig.String("web_base_url"), ctrl.Data["web_base_url"])
	}
}
