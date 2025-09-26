package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewMiddlewares_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "test"

	// when
	m := newMiddlewares(c, namespace)

	// then
	assert.NotNil(t, m)
	assert.Equal(t, c.RESTClient(), m.client)
	assert.Equal(t, namespace, m.ns)
}
