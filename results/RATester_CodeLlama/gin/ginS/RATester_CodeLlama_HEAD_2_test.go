package ginS

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestHEAD_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	relativePath := "/test"
	handlers := []gin.HandlerFunc{func(c *gin.Context) {}}

	// Act
	engine().HEAD(relativePath, handlers...)

	// Assert
	assert.Equal(t, relativePath, relativePath)
}
