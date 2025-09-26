package crd

import (
	"fmt"
	"testing"

	"k8s.io/client-go/rest"
)

func TestCreateClientFromConfig_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// arrange
	var c *rest.Config

	// act
	_, err := createClientFromConfig(c)

	// assert
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
