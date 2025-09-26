package ginS

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNoRoute_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new Gin router
	router := gin.New()

	// Define a handler function
	handler := func(c *gin.Context) {
		c.String(200, "Hello, world!")
	}

	// Call the method under test
	NoRoute(handler)

	// Create a new request
	req, _ := http.NewRequest("GET", "/non-existent-route", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Check the response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	if w.Body.String() != "Hello, world!" {
		t.Errorf("Expected body %s, but got %s", "Hello, world!", w.Body.String())
	}
}
