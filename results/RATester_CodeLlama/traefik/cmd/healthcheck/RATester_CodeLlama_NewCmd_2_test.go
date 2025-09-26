package healthcheck

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/paerser/cli"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestNewCmd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	traefikConfiguration := &static.Configuration{}
	loaders := []cli.ResourceLoader{}

	// when
	cmd := NewCmd(traefikConfiguration, loaders)

	// then
	require.NotNil(t, cmd)
	assert.Equal(t, "healthcheck", cmd.Name)
	assert.Equal(t, "Calls Traefik /ping endpoint (disabled by default) to check the health of Traefik.", cmd.Description)
	assert.Equal(t, traefikConfiguration, cmd.Configuration)
	assert.NotNil(t, cmd.Run)
	assert.Equal(t, loaders, cmd.Resources)
}
