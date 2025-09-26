package compress

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestConfigDefault_16(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var config []Config
	// Act
	var actual Config
	actual = configDefault(config...)
	// Assert
	assert.Equal(t, ConfigDefault, actual)
}
