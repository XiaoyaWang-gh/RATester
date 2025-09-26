package session

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestFresh_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	s := &Session{}
	s.fresh = true

	// Act
	actual := s.Fresh()

	// Assert
	assert.True(t, actual)
}
