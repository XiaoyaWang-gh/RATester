package render

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstance_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// Arrange
	var (
		r    HTMLDebug
		name string
		data any
	)

	// Act
	var actual Render = r.Instance(name, data)

	// Assert
	assert.NotNil(t, actual)
}
