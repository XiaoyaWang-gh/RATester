package web

import (
	"fmt"
	"testing"
)

func TestUnregisterFixedRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &HttpServer{
		Handlers: &ControllerRegister{
			routers: make(map[string]*Tree),
		},
	}

	app.Handlers.routers["GET"] = &Tree{
		prefix: "/test",
		leaves: []*leafInfo{
			{
				wildcards: []string{},
				runObject: "test",
			},
		},
	}

	app.UnregisterFixedRoute("/test", "GET")

	if _, ok := app.Handlers.routers["GET"]; ok {
		t.Error("Failed to unregister fixed route")
	}
}
