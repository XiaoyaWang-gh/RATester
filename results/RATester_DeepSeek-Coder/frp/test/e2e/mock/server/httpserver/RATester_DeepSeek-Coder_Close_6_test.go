package httpserver

import (
	"fmt"
	"net/http"
	"testing"
)

func TestClose_6(t *testing.T) {
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
		t.Errorf("Expected no error, got %v", err)
	}

	s.hs = nil
	err = s.Close()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
