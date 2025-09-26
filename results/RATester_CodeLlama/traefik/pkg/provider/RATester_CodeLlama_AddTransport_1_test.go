package provider

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestAddTransport_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configuration := &dynamic.HTTPConfiguration{
		ServersTransports: map[string]*dynamic.ServersTransport{},
	}

	transportName := "transportName"
	transport := &dynamic.ServersTransport{
		ServerName: "serverName",
	}

	if !AddTransport(configuration, transportName, transport) {
		t.Errorf("AddTransport() = false, want true")
	}

	if configuration.ServersTransports[transportName] == nil {
		t.Errorf("configuration.ServersTransports[transportName] = nil, want not nil")
	}

	if !reflect.DeepEqual(configuration.ServersTransports[transportName], transport) {
		t.Errorf("configuration.ServersTransports[transportName] = %v, want %v", configuration.ServersTransports[transportName], transport)
	}
}
