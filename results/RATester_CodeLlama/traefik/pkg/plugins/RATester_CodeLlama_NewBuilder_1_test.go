package plugins

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	client := &Client{}
	plugins := map[string]Descriptor{}
	localPlugins := map[string]LocalDescriptor{}

	// WHEN
	pb, err := NewBuilder(client, plugins, localPlugins)

	// THEN
	require.NoError(t, err)
	require.NotNil(t, pb)
}
