package hugolib

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestTestOptWarn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &IntegrationTestConfig{}
	TestOptWarn()(config)

	if config.LogLevel != logg.LevelWarn {
		t.Errorf("Expected LogLevel to be %v, but got %v", logg.LevelWarn, config.LogLevel)
	}
}
