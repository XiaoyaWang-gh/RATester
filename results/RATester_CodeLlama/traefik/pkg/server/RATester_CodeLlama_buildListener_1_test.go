package server

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestBuildListener_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	name := "test"
	config := &static.EntryPoint{
		Address: "127.0.0.1:8080",
	}

	listener, err := buildListener(ctx, name, config)
	require.NoError(t, err)
	require.NotNil(t, listener)
}
