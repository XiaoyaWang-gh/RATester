package ginS

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	gin.SetMode(gin.TestMode)
	router := gin.New()
	middleware1 := func(c *gin.Context) {
		// some logic
	}
	middleware2 := func(c *gin.Context) {
		// some logic
	}

	// Act
	router.Use(middleware1, middleware2)

	// Assert
	routes := router.Routes()
	if len(routes) != 1 {
		t.Errorf("Expected 1 route, got %d", len(routes))
	}
	if routes[0].Method != "ALL" {
		t.Errorf("Expected method to be 'ALL', got '%s'", routes[0].Method)
	}
	if routes[0].Path != "*" {
		t.Errorf("Expected path to be '*', got '%s'", routes[0].Path)
	}
}
