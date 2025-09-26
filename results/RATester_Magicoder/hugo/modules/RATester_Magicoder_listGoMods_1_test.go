package modules

import (
	"fmt"
	"testing"
)

func TestlistGoMods_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		GoModulesFilename: "go.mod",
		moduleConfig: Config{
			Imports: []Import{
				{Path: "github.com/example/module"},
			},
		},
	}

	modules, err := client.listGoMods()
	if err != nil {
		t.Errorf("listGoMods() error = %v", err)
		return
	}

	if len(modules) == 0 {
		t.Errorf("listGoMods() returned no modules")
	}

	for _, module := range modules {
		if module.Path != "github.com/example/module" {
			t.Errorf("listGoMods() returned unexpected module: %s", module.Path)
		}
	}
}
