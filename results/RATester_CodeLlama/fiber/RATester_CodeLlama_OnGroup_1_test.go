package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestOnGroup_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	handler := []func(Group) error{}
	// Act
	h.OnGroup(handler...)
	// Assert
	assert.Equal(t, handler, h.onGroup)
}
