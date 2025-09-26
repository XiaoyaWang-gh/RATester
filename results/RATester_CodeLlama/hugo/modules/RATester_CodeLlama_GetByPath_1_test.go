package modules

import (
	"fmt"
	"testing"
)

func TestGetByPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	modules := goModules{
		{
			Path: "github.com/go-modules/go-modules",
		},
		{
			Path: "github.com/go-modules/go-modules/v2",
		},
	}

	p := "github.com/go-modules/go-modules"
	m := modules.GetByPath(p)
	if m == nil {
		t.Errorf("expected module %s, got nil", p)
	}

	p = "github.com/go-modules/go-modules/v2"
	m = modules.GetByPath(p)
	if m == nil {
		t.Errorf("expected module %s, got nil", p)
	}

	p = "github.com/go-modules/go-modules/v3"
	m = modules.GetByPath(p)
	if m != nil {
		t.Errorf("expected nil, got module %s", p)
	}
}
