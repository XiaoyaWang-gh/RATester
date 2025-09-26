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

	Secret := "secret"
	name := "name"
	value := "value"
	others := []interface{}{"path", "/", "domain", "localhost", "secure", true, "httponly", true}

	ctrl.SetSecureCookie(Secret, name, value, others...)

	cookie, exists := ctrl.Ctx.GetSecureCookie(Secret, name)
	if !exists {
		t.Errorf("Expected cookie to exist, but it does not")
	}
	if cookie != value {
		t.Errorf("Expected cookie value to be %s, but got %s", value, cookie)
	}
}
