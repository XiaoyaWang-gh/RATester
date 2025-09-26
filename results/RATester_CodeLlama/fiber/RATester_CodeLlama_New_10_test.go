package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestNew_10(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{
		{
			ServerHeader: "test",
		},
	}
	// Act
	app := New(config...)
	// Assert
	assert.Equal(t, "test", app.config.ServerHeader)
}
