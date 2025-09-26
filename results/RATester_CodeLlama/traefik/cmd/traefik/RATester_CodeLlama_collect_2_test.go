package main

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestCollect_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	staticConfiguration := &static.Configuration{}
	// when
	collect(staticConfiguration)
	// then
	// TODO: add assertions
}
