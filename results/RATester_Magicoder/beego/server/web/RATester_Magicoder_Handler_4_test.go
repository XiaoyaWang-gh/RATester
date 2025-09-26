package web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &ControllerRegister{
		routers: make(map[string]*Tree),
	}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// your handler logic here
	})

	ctrl.Handler("/test", handler)

	// assert that the handler was added to the router
	if _, ok := ctrl.routers["/test"]; !ok {
		t.Error("Handler was not added to the router")
	}
}
