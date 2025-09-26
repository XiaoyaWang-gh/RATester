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

	// Test that the router is not nil
	if router == nil {
		t.Fatal("Router is nil")
	}

	// Test that the router returns a 404 for a non-existent route
	req, _ := http.NewRequest("GET", "/non-existent", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
