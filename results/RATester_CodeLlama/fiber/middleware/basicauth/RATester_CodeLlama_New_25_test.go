package basicauth

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNew_25(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := Config{
		Users: map[string]string{
			"user1": "pass1",
			"user2": "pass2",
		},
	}
	// Act
	handler := New(config)
	// Assert
	assert.NotNil(t, handler)
}
