package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestServersTransports_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "test"

	// when
	result := c.ServersTransports(namespace)

	// then
	assert.NotNil(t, result)
}
