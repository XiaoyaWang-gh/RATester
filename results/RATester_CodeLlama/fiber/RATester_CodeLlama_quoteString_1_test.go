package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestQuoteString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	app := &App{}
	raw := "hello"
	// Act
	quoted := app.quoteString(raw)
	// Assert
	assert.Equal(t, `"hello"`, quoted)
}
