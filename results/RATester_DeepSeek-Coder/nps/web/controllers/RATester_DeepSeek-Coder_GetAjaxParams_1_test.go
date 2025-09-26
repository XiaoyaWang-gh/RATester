package controllers

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"
)

func TestGetAjaxParams_1(t *testing.T) {
	type fields struct {
		Controller     beego.Controller
		controllerName string
		actionName     string
	}
	tests := []struct {
		name      string
		fields    fields
		wantStart int
		wantLimit int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &BaseController{
				Controller:     tt.fields.Controller,
				controllerName: tt.fields.controllerName,
				actionName:     tt.fields.actionName,
			}
			gotStart, gotLimit := s.GetAjaxParams()
			if gotStart != tt.wantStart {
				t.Errorf("GetAjaxParams() gotStart = %v, want %v", gotStart, tt.wantStart)
			}
			if gotLimit != tt.wantLimit {
				t.Errorf("GetAjaxParams() gotLimit = %v, want %v", gotLimit, tt.wantLimit)
			}
		})
	}
}
