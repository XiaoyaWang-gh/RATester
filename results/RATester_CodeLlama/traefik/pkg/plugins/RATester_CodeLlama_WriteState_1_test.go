package plugins

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{
		stateFile: "state.json",
	}

	plugins := map[string]Descriptor{
		"plugin1": {
			ModuleName: "plugin1",
			Version:    "v1.0.0",
		},
		"plugin2": {
			ModuleName: "plugin2",
			Version:    "v2.0.0",
		},
	}

	// when
	err := c.WriteState(plugins)

	// then
	require.NoError(t, err)

	// when
	b, err := os.ReadFile(c.stateFile)

	// then
	require.NoError(t, err)

	// when
	var m map[string]string
	err = json.Unmarshal(b, &m)

	// then
	require.NoError(t, err)

	// when
	require.Equal(t, "v1.0.0", m["plugin1"])
	require.Equal(t, "v2.0.0", m["plugin2"])
}
