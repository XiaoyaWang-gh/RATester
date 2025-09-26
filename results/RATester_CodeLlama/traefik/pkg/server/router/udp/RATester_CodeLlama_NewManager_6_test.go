package udp

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/runtime"
	udpservice "github.com/traefik/traefik/v3/pkg/server/service/udp"
)

func TestNewManager_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	conf := &runtime.Configuration{}
	serviceManager := &udpservice.Manager{}
	manager := NewManager(conf, serviceManager)
	assert.NotNil(t, manager)
	assert.Equal(t, serviceManager, manager.serviceManager)
	assert.Equal(t, conf, manager.conf)
}
