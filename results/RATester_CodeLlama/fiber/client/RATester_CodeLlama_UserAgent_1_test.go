package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestUserAgent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	r := &Request{}
	r.userAgent = "test"
	// Act
	actual := r.UserAgent()
	// Assert
	assert.Equal(t, "test", actual)
}
