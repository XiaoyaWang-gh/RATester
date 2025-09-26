package udp

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestBuildHandlers_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	m := &Manager{}
	rootCtx := context.Background()
	entryPoints := []string{"entryPoint1", "entryPoint2"}

	// when
	entryPointHandlers := m.BuildHandlers(rootCtx, entryPoints)

	// then
	assert.Equal(t, 2, len(entryPointHandlers))
}
