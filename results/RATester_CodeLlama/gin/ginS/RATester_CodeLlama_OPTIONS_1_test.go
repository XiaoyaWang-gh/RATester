package ginS

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestOPTIONS_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	relativePath := "test"
	handlers := []gin.HandlerFunc{}
	expected := engine().OPTIONS(relativePath, handlers...)

	// Act
	actual := OPTIONS(relativePath, handlers...)

	// Assert
	assert.Equal(t, expected, actual)
}
