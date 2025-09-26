package modules

import (
	"fmt"
	"testing"
)

func TestGetMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	modules := goModules{
		{
			Path: "github.com/golang/mock",
			Main: false,
		},
		{
			Path: "github.com/golang/mock",
			Main: true,
		},
		{
			Path: "github.com/golang/mock",
			Main: false,
		},
	}

	if modules.GetMain() != modules[1] {
		t.Errorf("expected %v, got %v", modules[1], modules.GetMain())
	}
}
