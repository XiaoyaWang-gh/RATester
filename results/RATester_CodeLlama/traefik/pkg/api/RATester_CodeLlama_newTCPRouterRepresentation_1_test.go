package api

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewTCPRouterRepresentation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	name := "test"
	rt := &runtime.TCPRouterInfo{}
	// when
	result := newTCPRouterRepresentation(name, rt)
	// then
	assert.Equal(t, name, result.Name)
	assert.Equal(t, getProviderName(name), result.Provider)
}
