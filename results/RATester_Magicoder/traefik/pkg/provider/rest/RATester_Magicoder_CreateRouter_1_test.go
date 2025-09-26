package rest

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestCreateRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	provider := &Provider{
		Insecure:          true,
		configurationChan: make(chan dynamic.Message),
	}

	router := provider.CreateRouter()

	if router == nil {
		t.Error("Expected router to be created, but got nil")
	}

	route := router.Get("/api/providers/{provider}")
	if route == nil {
		t.Error("Expected route to be created, but got nil")
	}

	handler := route.GetHandler()
	if handler == nil {
		t.Error("Expected handler to be set, but got nil")
	}

	if handler != provider {
		t.Error("Expected handler to be the provider, but got a different handler")
	}
}
