package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestaddToPolicy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		policies: make(map[string]*Tree),
	}

	ctrl.addToPolicy("GET", "/test", func(ctx *context.Context) {
		// Add your test logic here
	})

	if _, ok := ctrl.policies["GET"]; !ok {
		t.Error("Expected policy to be added")
	}
}
