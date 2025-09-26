package helmet

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			XSSProtection: "1",
		},
	}

	// Act
	handler := New(config...)

	// Assert
	assert.NotNil(t, handler)
}
