package metrics

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestHasService_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dynamicConfig{
		services: map[string]map[string]bool{
			"service1": {
				"server1": true,
				"server2": true,
			},
			"service2": {
				"server3": true,
				"server4": true,
			},
		},
	}
	serviceName := "service1"
	assert.True(t, d.hasService(serviceName))
	serviceName = "service3"
	assert.False(t, d.hasService(serviceName))
}
