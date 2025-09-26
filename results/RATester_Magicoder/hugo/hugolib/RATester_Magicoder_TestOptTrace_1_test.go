package hugolib

import (
	"fmt"
	"testing"

	"github.com/bep/logg"
)

func TestTestOptTrace_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := &IntegrationTestConfig{}
	opt := TestOptTrace()
	opt(config)

	if config.LogLevel != logg.LevelTrace {
		t.Errorf("Expected LogLevel to be %s, but got %s", logg.LevelTrace, config.LogLevel)
	}
}
