package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewServersTransports_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "test"

	// when
	serversTransports := newServersTransports(c, namespace)

	// then
	assert.NotNil(t, serversTransports)
	assert.Equal(t, c.RESTClient(), serversTransports.client)
	assert.Equal(t, namespace, serversTransports.ns)
}
