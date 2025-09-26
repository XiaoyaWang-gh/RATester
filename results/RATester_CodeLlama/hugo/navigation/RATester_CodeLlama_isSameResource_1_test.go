package navigation

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsSameResource_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	m := &MenuEntry{
		MenuConfig: MenuConfig{
			PageRef: "page1",
		},
	}
	inme := &MenuEntry{
		MenuConfig: MenuConfig{
			PageRef: "page1",
		},
	}

	// when
	result := m.isSameResource(inme)

	// then
	assert.True(t, result)
}
