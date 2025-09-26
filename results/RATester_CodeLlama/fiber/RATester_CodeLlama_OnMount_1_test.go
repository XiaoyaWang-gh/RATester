package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestOnMount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	handler := []func(*App) error{}
	// Act
	h.OnMount(handler...)
	// Assert
	assert.Equal(t, handler, h.onMount)
}
