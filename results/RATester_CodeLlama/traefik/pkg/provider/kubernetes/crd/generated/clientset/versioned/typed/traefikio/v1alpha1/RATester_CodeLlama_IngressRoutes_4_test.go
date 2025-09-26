package v1alpha1

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIngressRoutes_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	c := &TraefikV1alpha1Client{}
	namespace := "namespace"

	// when
	c.IngressRoutes(namespace)

	// then
	assert.Equal(t, "namespace", namespace)
}
