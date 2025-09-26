package identity

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewFindFirstManagerIdentityProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var m Manager
	var id Identity

	// Act
	var actual FindFirstManagerIdentityProvider
	actual = NewFindFirstManagerIdentityProvider(m, id)

	// Assert
	assert.NotNil(t, actual)
}
