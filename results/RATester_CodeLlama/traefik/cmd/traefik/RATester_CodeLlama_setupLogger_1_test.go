package main

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestSetupLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	staticConfiguration := &static.Configuration{}

	// when
	setupLogger(staticConfiguration)

	// then
	// TODO
}
