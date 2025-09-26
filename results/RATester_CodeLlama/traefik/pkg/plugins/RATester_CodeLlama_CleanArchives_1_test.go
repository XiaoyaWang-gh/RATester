package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCleanArchives_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &Client{
		archives: "archives",
	}
	plugins := map[string]Descriptor{
		"plugin1": {
			ModuleName: "plugin1",
			Version:    "1.0.0",
		},
		"plugin2": {
			ModuleName: "plugin2",
			Version:    "1.0.0",
		},
	}

	// when
	err := c.CleanArchives(plugins)

	// then
	require.NoError(t, err)
}
