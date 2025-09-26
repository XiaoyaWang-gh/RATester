package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestOnFork_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	handler := []func(int) error{}
	h.OnFork(handler...)
	// Act
	h.executeOnForkHooks(1)
	// Assert
	assert.Equal(t, handler, h.onFork)
}
