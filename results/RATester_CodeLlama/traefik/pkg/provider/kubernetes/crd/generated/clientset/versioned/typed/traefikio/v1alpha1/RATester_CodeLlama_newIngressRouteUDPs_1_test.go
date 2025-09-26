package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewIngressRouteUDPs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "test"

	// when
	actual := newIngressRouteUDPs(c, namespace)

	// then
	assert.NotNil(t, actual)
	assert.Equal(t, c.RESTClient(), actual.client)
	assert.Equal(t, namespace, actual.ns)
}
