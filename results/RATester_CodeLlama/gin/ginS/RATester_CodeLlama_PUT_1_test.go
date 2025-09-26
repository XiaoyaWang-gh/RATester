package ginS

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestPUT_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	relativePath := "/test"
	handlers := []gin.HandlerFunc{}
	expected := engine().PUT(relativePath, handlers...)

	// Act
	actual := PUT(relativePath, handlers...)

	// Assert
	assert.Equal(t, expected, actual)
}
