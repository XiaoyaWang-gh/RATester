package web

import (
	"fmt"
	"testing"
)

func TestRESTRouter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctrl := &Controller{}
	rootpath := "/test"
	server := RESTRouter(rootpath, ctrl)

	if server == nil {
		t.Error("Expected server to not be nil")
	}

	if server.Handlers == nil {
		t.Error("Expected server.Handlers to not be nil")
	}

	if server.Server == nil {
		t.Error("Expected server.Server to not be nil")
	}

	if server.Cfg == nil {
		t.Error("Expected server.Cfg to not be nil")
	}

	if len(server.LifeCycleCallbacks) != 0 {
		t.Error("Expected server.LifeCycleCallbacks to be empty")
	}
}
