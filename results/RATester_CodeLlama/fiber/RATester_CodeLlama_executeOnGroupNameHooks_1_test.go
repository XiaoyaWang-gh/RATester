package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestExecuteOnGroupNameHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var group Group
	var h Hooks
	h.onGroupName = []OnGroupNameHandler{
		func(group Group) error {
			group.Prefix = "test"
			return nil
		},
	}
	// when
	err := h.executeOnGroupNameHooks(group)
	// then
	assert.NoError(t, err)
	assert.Equal(t, "test", group.Prefix)
}
