package fiber

import (
	"fmt"
	"testing"
)

func TestConfigDependentPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{
		config: Config{
			UnescapePath: true,
		},
		getString: func(b []byte) string {
			return string(b)
		},
	}

	ctx := &DefaultCtx{
		app:                 app,
		pathOriginal:        "/test/path",
		pathBuffer:          []byte("/test/path"),
		detectionPathBuffer: []byte("/test/path"),
	}

	ctx.configDependentPaths()

	if ctx.path != "/test/path" {
		t.Errorf("Expected path to be '/test/path', got %s", ctx.path)
	}

	if ctx.detectionPath != "/test/path" {
		t.Errorf("Expected detectionPath to be '/test/path', got %s", ctx.detectionPath)
	}

	if ctx.treePath != "/tes" {
		t.Errorf("Expected treePath to be '/tes', got %s", ctx.treePath)
	}
}
