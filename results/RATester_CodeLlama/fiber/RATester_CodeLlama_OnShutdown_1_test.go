package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestOnShutdown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	handler := func() error {
		return nil
	}
	// Act
	h.OnShutdown(handler)
	// Assert
	assert.Equal(t, handler, h.onShutdown[0])
}
