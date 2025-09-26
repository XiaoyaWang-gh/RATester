package fiber

import (
	"fmt"
	"testing"
)

func TestJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var data any
	var ctype []string
	var c *DefaultCtx
	// Act
	err := c.JSON(data, ctype...)
	// Assert
	if err != nil {
		t.Errorf("JSON returned an error: %v", err)
	}
}
