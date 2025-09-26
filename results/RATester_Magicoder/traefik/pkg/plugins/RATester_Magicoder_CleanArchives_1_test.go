package plugins

import (
	"fmt"
	"testing"
)

func TestCleanArchives_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		stateFile: "testdata/state.json",
	}

	plugins := map[string]Descriptor{
		"plugin1": {ModuleName: "plugin1", Version: "1.0.0"},
		"plugin2": {ModuleName: "plugin2", Version: "2.0.0"},
	}

	err := client.CleanArchives(plugins)
	if err != nil {
		t.Errorf("CleanArchives() error = %v", err)
	}

	// Add assertions here to check if the archives were cleaned as expected
}
