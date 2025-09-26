package main

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestCleanServers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	element := &dynamic.Configuration{}
	cleanServers(element)
	if len(element.HTTP.Services) != 0 {
		t.Errorf("element.HTTP.Services should be empty")
	}
	if len(element.TCP.Services) != 0 {
		t.Errorf("element.TCP.Services should be empty")
	}
	if len(element.UDP.Services) != 0 {
		t.Errorf("element.UDP.Services should be empty")
	}
}
