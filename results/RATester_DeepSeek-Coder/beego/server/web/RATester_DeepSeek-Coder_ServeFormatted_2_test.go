package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestServeFormatted_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
		Data: map[interface{}]interface{}{
			"key": "value",
		},
	}

	err := ctrl.ServeFormatted()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	err = ctrl.ServeFormatted(true)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
