package proxy

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/server/controller"
)

func TestGetResourceController_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pxy := &BaseProxy{
		rc: &controller.ResourceController{},
	}

	rc := pxy.GetResourceController()

	if rc != pxy.rc {
		t.Errorf("Expected %v, got %v", pxy.rc, rc)
	}
}
