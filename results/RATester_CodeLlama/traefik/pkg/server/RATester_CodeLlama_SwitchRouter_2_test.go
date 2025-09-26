package server

import (
	"fmt"
	"testing"

	tcprouter "github.com/traefik/traefik/v3/pkg/server/router/tcp"
)

func TestSwitchRouter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	e := &TCPEntryPoint{}
	rt := &tcprouter.Router{}

	// when
	e.SwitchRouter(rt)

	// then
	// TODO
}
