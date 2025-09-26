package testhelpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestWithServiceName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	serviceName := "my-service"
	router := &dynamic.Router{}
	WithServiceName(serviceName)(router)
	assert.Equal(t, serviceName, router.Service)
}
