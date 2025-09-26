package fiber

import (
	"fmt"
	"testing"
)

func TestconfigDependentPaths_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &DefaultCtx{
		app: &App{
			config: Config{
				UnescapePath: true,
			},
		},
		pathOriginal: "test/path",
	}

	ctx.configDependentPaths()

	if ctx.path != "test/path" {
		t.Errorf("Expected path to be 'test/path', but got '%s'", ctx.path)
	}

	if ctx.detectionPath != "test/path" {
		t.Errorf("Expected detectionPath to be 'test/path', but got '%s'", ctx.detectionPath)
	}

	if ctx.treePath != "test/" {
		t.Errorf("Expected treePath to be 'test/', but got '%s'", ctx.treePath)
	}
}
