package visitor

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewKCPTunnelSession_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	// Act
	result := NewKCPTunnelSession()
	// Assert
	assert.NotNil(t, result)
}
