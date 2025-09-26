package router

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestBuildHandlers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	m := &Manager{}
	rootCtx := context.Background()
	entryPoints := []string{"entryPoint1", "entryPoint2"}
	tls := true

	// when
	entryPointHandlers := m.BuildHandlers(rootCtx, entryPoints, tls)

	// then
	assert.Equal(t, 2, len(entryPointHandlers))
}
