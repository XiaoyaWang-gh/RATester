package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildDefaultHTTPRouter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	router := BuildDefaultHTTPRouter()

	if router == nil {
		t.Errorf("Expected a non-nil router, got nil")
	}

	server := httptest.NewServer(router)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %v, got %v", http.StatusNotFound, resp.StatusCode)
	}
}
