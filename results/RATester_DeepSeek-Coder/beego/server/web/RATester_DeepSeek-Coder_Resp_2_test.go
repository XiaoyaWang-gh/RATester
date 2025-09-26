package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestResp_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	data := "test data"
	err := ctrl.Resp(data)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
