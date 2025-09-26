package web

import (
	"fmt"
	"testing"

	beecontext "github.com/beego/beego/v2/server/web/context"
)

func TestinsertFilterRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		enableFilter: false,
		filters:      [5][]*FilterRouter{},
	}

	testFilter := &FilterRouter{
		filterFunc: func(ctx *beecontext.Context) {
			// do something
		},
	}

	err := ctrl.insertFilterRouter(0, testFilter)
	if err != nil {
		t.Errorf("insertFilterRouter failed, err: %v", err)
	}

	if !ctrl.enableFilter {
		t.Errorf("enableFilter should be true after insertFilterRouter")
	}

	if len(ctrl.filters[0]) != 1 {
		t.Errorf("filters[0] should have 1 filter after insertFilterRouter")
	}

	if ctrl.filters[0][0] != testFilter {
		t.Errorf("filters[0][0] should be the testFilter after insertFilterRouter")
	}
}
