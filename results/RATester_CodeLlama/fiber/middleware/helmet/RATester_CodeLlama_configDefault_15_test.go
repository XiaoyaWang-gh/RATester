package helmet

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestConfigDefault_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var config Config
	var expected Config

	// Act
	actual := configDefault(config)

	// Assert
	assert.Equal(t, expected, actual)
}
