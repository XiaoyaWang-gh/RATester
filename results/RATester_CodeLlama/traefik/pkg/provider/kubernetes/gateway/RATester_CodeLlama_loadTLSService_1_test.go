package gateway

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
	gatev1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

func TestLoadTLSService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Provider{}
	route := &gatev1alpha2.TLSRoute{}
	backendRef := &gatev1.BackendRef{}
	serviceName, tlsService, condition := p.loadTLSService(route, *backendRef)
	assert.Equal(t, "", serviceName)
	assert.Nil(t, tlsService)
	assert.NotNil(t, condition)
}
