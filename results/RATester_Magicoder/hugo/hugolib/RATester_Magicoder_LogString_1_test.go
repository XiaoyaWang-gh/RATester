package hugolib

import (
	"fmt"
	"testing"
)

func TestLogString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &IntegrationTestBuilder{
		lastBuildLog: "test log",
	}

	log := b.LogString()

	if log != "test log" {
		t.Errorf("Expected 'test log', got '%s'", log)
	}
}
