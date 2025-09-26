package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestServeJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
		Data: map[interface{}]interface{}{
			"json": "test",
		},
	}

	err := ctrl.ServeJSON()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = ctrl.ServeJSON(true)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
