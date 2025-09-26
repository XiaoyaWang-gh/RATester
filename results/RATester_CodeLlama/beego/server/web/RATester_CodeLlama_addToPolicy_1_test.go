package web

import (
	"fmt"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestAddToPolicy_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &ControllerRegister{}
	p.addToPolicy("GET", "/", func(ctx *context.Context) {})
	if len(p.policies) != 1 {
		t.Fatalf("TestaddToPolicy failed")
	}
}
