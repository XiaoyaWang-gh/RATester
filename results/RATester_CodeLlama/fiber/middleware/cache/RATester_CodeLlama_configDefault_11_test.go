package cache

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestConfigDefault_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := []Config{}

	// Act
	actual := configDefault(config...)

	// Assert
	assert.Equal(t, ConfigDefault, actual)
}
