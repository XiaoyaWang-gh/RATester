package rest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/safe"
)

func TestProvide_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	configurationChan := make(chan dynamic.Message, 1)
	pool := &safe.Pool{}
	err := p.Provide(configurationChan, pool)
	require.NoError(t, err)
	require.NotNil(t, p.configurationChan)
}
