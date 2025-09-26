package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSetXMLUnmarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	c := &Client{}
	f := func(data []byte, v any) error {
		return nil
	}
	// Act
	c.SetXMLUnmarshal(f)
	// Assert
	assert.Equal(t, f, c.xmlUnmarshal)
}
