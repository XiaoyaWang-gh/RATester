package http

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClose_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	server := &Server{
		hs: &http.Server{},
	}

	err := server.Close()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
