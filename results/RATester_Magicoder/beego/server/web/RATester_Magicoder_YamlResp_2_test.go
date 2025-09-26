package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestYamlResp_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	data := map[string]interface{}{
		"key": "value",
	}

	err := ctrl.YamlResp(data)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
