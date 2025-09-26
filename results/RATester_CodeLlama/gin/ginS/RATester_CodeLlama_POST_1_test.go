package ginS

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
)

func TestPOST_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	relativePath := "/test"
	handlers := []gin.HandlerFunc{func(c *gin.Context) {}}
	expected := engine().POST(relativePath, handlers...)

	// Act
	actual := POST(relativePath, handlers...)

	// Assert
	assert.Equal(t, expected, actual)
}
