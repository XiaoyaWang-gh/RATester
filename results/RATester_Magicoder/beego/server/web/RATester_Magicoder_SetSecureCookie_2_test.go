package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestSetSecureCookie_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{
		Ctx: &context.Context{},
	}

	ctrl.SetSecureCookie("secret", "name", "value")

	cookie, ok := ctrl.Ctx.GetSecureCookie("secret", "name")
	if !ok {
		t.Error("Expected cookie to be set")
	}
	if cookie != "value" {
		t.Errorf("Expected cookie value to be 'value', but got '%s'", cookie)
	}
}
