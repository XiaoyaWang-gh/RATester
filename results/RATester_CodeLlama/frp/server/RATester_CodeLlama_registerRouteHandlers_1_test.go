package server

import (
	"fmt"
	"testing"

	httppkg "github.com/fatedier/frp/pkg/util/http"
)

func TestRegisterRouteHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	helper := &httppkg.RouterRegisterHelper{}
	svr.registerRouteHandlers(helper)

	// TODO: check the router
}
