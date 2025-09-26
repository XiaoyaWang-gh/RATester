package hugolib

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestTestOptDebug_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &IntegrationTestConfig{
		LogLevel: logg.LevelInfo,
	}

	TestOptDebug()(config)

	if config.LogLevel != logg.LevelDebug {
		t.Errorf("Expected LogLevel to be %v, but got %v", logg.LevelDebug, config.LogLevel)
	}
}
