package ecs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetServiceName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	instance := ecsInstance{
		Name: "test",
	}
	// when
	result := getServiceName(instance)
	// then
	assert.Equal(t, "test", result)
}
