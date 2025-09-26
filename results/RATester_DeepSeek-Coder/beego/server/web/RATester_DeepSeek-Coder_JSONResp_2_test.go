package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestJSONResp_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	data := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John Doe",
		Age:  30,
	}

	err := ctrl.JSONResp(data)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
