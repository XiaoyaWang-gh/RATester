package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestCond_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{
		prefix:   "/test",
		handlers: &ControllerRegister{},
	}

	cond := func(ctx *beecontext.Context) bool {
		return true
	}

	n.Cond(cond)

	// Assert that the condition was added to the namespace's handlers
	if len(n.handlers.filters[BeforeRouter]) == 0 {
		t.Error("Expected condition to be added to handlers")
	}
}
