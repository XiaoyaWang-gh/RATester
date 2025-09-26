package api

import (
	"fmt"
	"strings"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/runtime"
)

func TestNewServiceRepresentation_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	name := "name"
	si := &runtime.ServiceInfo{}
	// when
	result := newServiceRepresentation(name, si)
	// then
	assert.Equal(t, si, result.ServiceInfo)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, getProviderName(name), result.Provider)
	assert.Equal(t, si.GetAllStatus(), result.ServerStatus)
	assert.Equal(t, strings.ToLower(extractType(si.Service)), result.Type)
}
