package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestOnGroupName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	h := &Hooks{}
	handler := []func(Group) error{}
	h.OnGroupName(handler...)
	// Act
	h.executeOnGroupNameHooks(Group{})
	// Assert
	assert.Equal(t, handler, h.onGroupName)
}
