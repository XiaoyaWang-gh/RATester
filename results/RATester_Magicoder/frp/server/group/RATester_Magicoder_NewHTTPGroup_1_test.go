package group

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
)

func TestNewHTTPGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctl := &HTTPGroupController{
		groups:      make(map[string]*HTTPGroup),
		vhostRouter: &vhost.Routers{},
	}

	group := NewHTTPGroup(ctl)

	if group == nil {
		t.Error("NewHTTPGroup returned nil")
	}

	if group.createFuncs == nil {
		t.Error("NewHTTPGroup.createFuncs is nil")
	}

	if group.pxyNames == nil {
		t.Error("NewHTTPGroup.pxyNames is nil")
	}

	if group.ctl != ctl {
		t.Error("NewHTTPGroup.ctl is not set correctly")
	}
}
