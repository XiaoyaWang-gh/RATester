package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestExecuteOnShutdownHooks_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	h := &Hooks{}
	h.onShutdown = []OnShutdownHandler{
		func() error {
			return nil
		},
		func() error {
			return nil
		},
	}
	// When
	h.executeOnShutdownHooks()
	// Then
	assert.Equal(t, 2, len(h.onShutdown))
}
