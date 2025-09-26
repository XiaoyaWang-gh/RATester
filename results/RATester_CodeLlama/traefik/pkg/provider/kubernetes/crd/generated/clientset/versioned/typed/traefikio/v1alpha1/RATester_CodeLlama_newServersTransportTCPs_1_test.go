package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewServersTransportTCPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "test"

	// when
	serversTransportTCPs := newServersTransportTCPs(c, namespace)

	// then
	assert.NotNil(t, serversTransportTCPs)
	assert.Equal(t, c.RESTClient(), serversTransportTCPs.client)
	assert.Equal(t, namespace, serversTransportTCPs.ns)
}
