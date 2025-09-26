package net

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClose_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	listener := &WebsocketListener{
		server: &http.Server{},
	}

	err := listener.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
