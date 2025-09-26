package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestInsertFilterRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrlRegister := &ControllerRegister{
		enableFilter: false,
		filters:      [FinishRouter + 1][]*FilterRouter{},
	}

	filterRouter := &FilterRouter{
		filterFunc: func(ctx *beecontext.Context) {
			// do something
		},
	}

	err := ctrlRegister.insertFilterRouter(BeforeStatic, filterRouter)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !ctrlRegister.enableFilter {
		t.Errorf("Expected enableFilter to be true, got false")
	}

	if len(ctrlRegister.filters[BeforeStatic]) != 1 {
		t.Errorf("Expected 1 filter at BeforeStatic position, got %d", len(ctrlRegister.filters[BeforeStatic]))
	}

	if ctrlRegister.filters[BeforeStatic][0] != filterRouter {
		t.Errorf("Expected filterRouter at BeforeStatic position, got different filter")
	}
}
