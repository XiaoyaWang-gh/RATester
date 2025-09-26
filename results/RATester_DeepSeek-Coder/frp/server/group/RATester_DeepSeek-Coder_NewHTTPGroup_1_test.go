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

	httpGroup := NewHTTPGroup(ctl)

	if httpGroup == nil {
		t.Errorf("NewHTTPGroup() = %v, want not nil", httpGroup)
	}

	if httpGroup.createFuncs == nil {
		t.Errorf("NewHTTPGroup().createFuncs = %v, want not nil", httpGroup.createFuncs)
	}

	if httpGroup.pxyNames == nil {
		t.Errorf("NewHTTPGroup().pxyNames = %v, want not nil", httpGroup.pxyNames)
	}

	if httpGroup.ctl != ctl {
		t.Errorf("NewHTTPGroup().ctl = %v, want %v", httpGroup.ctl, ctl)
	}
}
