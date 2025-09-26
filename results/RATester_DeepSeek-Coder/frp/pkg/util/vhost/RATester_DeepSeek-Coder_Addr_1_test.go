package vhost

import (
	"fmt"
	"testing"
)

func TestAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &Listener{
		name:     "test",
		location: "localhost:8080",
	}

	addr := l.Addr()
	if addr != nil {
		t.Errorf("Expected nil, got %v", addr)
	}
}
