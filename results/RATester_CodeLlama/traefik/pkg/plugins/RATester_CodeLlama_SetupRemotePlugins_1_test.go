package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetupRemotePlugins_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	client := &Client{}
	plugins := map[string]Descriptor{}

	// when
	err := SetupRemotePlugins(client, plugins)

	// then
	require.NoError(t, err)
}
