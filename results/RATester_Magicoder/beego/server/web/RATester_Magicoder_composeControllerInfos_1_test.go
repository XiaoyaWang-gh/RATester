package web

import (
	"fmt"
	"testing"
)

func TestcomposeControllerInfos_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	tree := &Tree{}
	routerInfos := []*ControllerInfo{}

	composeControllerInfos(tree, &routerInfos)

	if len(routerInfos) != 0 {
		t.Errorf("Expected 0 routerInfos, but got %d", len(routerInfos))
	}

	controllerInfo := &ControllerInfo{
		pattern: "test",
	}
	tree.leaves = []*leafInfo{{runObject: controllerInfo}}

	composeControllerInfos(tree, &routerInfos)

	if len(routerInfos) != 1 {
		t.Errorf("Expected 1 routerInfos, but got %d", len(routerInfos))
	}

	if routerInfos[0] != controllerInfo {
		t.Errorf("Expected %v, but got %v", controllerInfo, routerInfos[0])
	}
}
