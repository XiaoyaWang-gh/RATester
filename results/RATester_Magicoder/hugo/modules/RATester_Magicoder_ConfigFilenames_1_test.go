package modules

import (
	"fmt"
	"testing"
)

func TestConfigFilenames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	adapter := &moduleAdapter{
		configFilenames: []string{"config1.yaml", "config2.yaml"},
	}

	filenames := adapter.ConfigFilenames()

	if len(filenames) != 2 {
		t.Errorf("Expected 2 filenames, got %d", len(filenames))
	}

	if filenames[0] != "config1.yaml" || filenames[1] != "config2.yaml" {
		t.Errorf("Unexpected filenames: %v", filenames)
	}
}
