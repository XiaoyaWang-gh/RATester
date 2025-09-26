package babel

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/resources"
)

func TestNew_47(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	var rs *resources.Spec

	// Act
	var client *Client
	client = New(rs)

	// Assert
	assert.NotNil(t, client)
}
