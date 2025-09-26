package vhost

import (
	"fmt"
	"testing"
)

func TestAdd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Routers{
		indexByDomain: make(map[string]routerByHTTPUser),
	}

	err := r.Add("example.com", "/", "user1", "payload1")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	vr, exist := r.Get("example.com", "/", "user1")
	if !exist {
		t.Fatalf("Expected router to exist")
	}

	if vr.payload != "payload1" {
		t.Fatalf("Expected payload to be 'payload1', but got %v", vr.payload)
	}

	err = r.Add("example.com", "/", "user1", "payload2")
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if err != ErrRouterConfigConflict {
		t.Fatalf("Expected error to be ErrRouterConfigConflict, but got %v", err)
	}
}
