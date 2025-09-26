package group

import (
	"fmt"
	"testing"

	"github.com/fatedier/frp/pkg/util/vhost"
	"github.com/zeebo/assert"
)

func TestNewHTTPGroupController_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	vhostRouter := &vhost.Routers{}
	ctl := NewHTTPGroupController(vhostRouter)
	assert.NotNil(t, ctl)
}
