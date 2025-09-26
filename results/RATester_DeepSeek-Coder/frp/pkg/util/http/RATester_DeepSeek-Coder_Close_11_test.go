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

	s := &Server{
		hs: &http.Server{},
	}

	err := s.Close()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
