package circuitbreaker

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeHTTP_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var rw http.ResponseWriter
	var req *http.Request
	var c *circuitBreaker

	// Act
	c.ServeHTTP(rw, req)

	// Assert
	// ...
}
