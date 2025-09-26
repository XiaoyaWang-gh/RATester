package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestCreatePluginBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	staticConfiguration := &static.Configuration{}

	// when
	builder, err := createPluginBuilder(staticConfiguration)

	// then
	require.NoError(t, err)
	require.NotNil(t, builder)
}
