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

	// Arrange
	relativePath := "/test"
	handlers := []gin.HandlerFunc{func(c *gin.Context) {}}

	// Act
	GET(relativePath, handlers...)

	// Assert
	// TODO
}
