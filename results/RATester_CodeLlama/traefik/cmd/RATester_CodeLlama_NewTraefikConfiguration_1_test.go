package cmd

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewTraefikConfiguration_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	expected := &TraefikCmdConfiguration{
		Configuration: static.Configuration{
			Global: &static.Global{
				CheckNewVersion: true,
			},
			EntryPoints: make(static.EntryPoints),
			Providers: &static.Providers{
				ProvidersThrottleDuration: ptypes.Duration(2 * time.Second),
			},
			ServersTransport: &static.ServersTransport{
				MaxIdleConnsPerHost: 200,
			},
			TCPServersTransport: &static.TCPServersTransport{
				DialTimeout:   ptypes.Duration(30 * time.Second),
				DialKeepAlive: ptypes.Duration(15 * time.Second),
			},
		},
		ConfigFile: "",
	}

	// Act
	actual := NewTraefikConfiguration()

	// Assert
	assert.Equal(t, expected, actual)
}
