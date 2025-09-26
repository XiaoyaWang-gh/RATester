package ginS

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGET_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	relativePath := "/test"
	handlers := []gin.HandlerFunc{
		func(c *gin.Context) {
			// some handler logic
		},
	}

	result := GET(relativePath, handlers...)

	if result == nil {
		t.Error("Expected a non-nil result, but got nil")
	}
}
