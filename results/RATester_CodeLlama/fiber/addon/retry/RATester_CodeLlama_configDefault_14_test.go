package retry

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestConfigDefault_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var config []Config
	// Act
	actual := configDefault(config...)
	// Assert
	assert.Equal(t, DefaultConfig, actual)
}
