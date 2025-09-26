package session

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestCfgSessionIdLength_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	length := int64(10)
	expected := ManagerConfig{
		SessionIDLength: length,
	}
	// Act
	actual := CfgSessionIdLength(length)
	// Assert
	assert.Equal(t, expected, actual)
}
