package plugins

import (
	"fmt"
	"testing"
)

func TestNewWasmMiddlewareBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	goPath := "/path/to/go"
	moduleName := "module"
	wasmPath := "/path/to/wasm"
	settings := Settings{
		Envs:   []string{"env1=value1", "env2=value2"},
		Mounts: []string{"/path/to/mount1", "/path/to/mount2"},
	}

	_, err := newWasmMiddlewareBuilder(goPath, moduleName, wasmPath, settings)
	if err != nil {
		t.Fatalf("newWasmMiddlewareBuilder() failed: %v", err)
	}
}
