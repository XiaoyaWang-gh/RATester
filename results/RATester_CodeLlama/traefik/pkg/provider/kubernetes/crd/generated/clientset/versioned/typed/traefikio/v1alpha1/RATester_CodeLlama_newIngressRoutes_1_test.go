package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/stretchr/testify/require"
)

func TestNewIngressRoutes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "test"

	// when
	result := newIngressRoutes(c, namespace)

	// then
	require.NotNil(t, result)
	assert.Equal(t, c.RESTClient(), result.client)
	assert.Equal(t, namespace, result.ns)
}
